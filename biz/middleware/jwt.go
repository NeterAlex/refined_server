package middleware

import (
	"Refined_service/biz/dal/sqlite"
	"Refined_service/biz/model/user"
	"Refined_service/biz/pack"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/jwt"
)

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}

type User struct {
	Username string `json:"username,omitempty"`
	ID       string `json:"userID,omitempty"`
	Class    string `json:"status,omitempty"`
}

var Username = "username"
var Uid = "uid"
var Status = "status"

func AuthInit(secretKey string) *jwt.HertzJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "refined",
		SigningAlgorithm: "HS256",
		Key:              []byte(secretKey),
		Timeout:          time.Hour * 72,
		MaxRefresh:       time.Hour * 72,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			//Check if username matches password
			if pack.CheckAuthValid(username, password) {
				u, _, err := sqlite.Query[user.User]("username = ?", username)
				if err != nil {
					return User{}, nil
				}
				return &User{
					Username: username,
					ID:       strconv.FormatInt(u[0].ID, 10),
					Class:    u[0].Status,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					Username: v.Username,
					Uid:      v.ID,
					Status:   v.Class,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": http.StatusOK,
			})
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				Username: claims[Username].(string),
				ID:       claims[Uid].(string),
			}
		},
		IdentityKey: Username,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName:               "Bearer",
		WithoutDefaultTokenHeadName: false,
		TimeFunc:                    time.Now,
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
		SendCookie:        true,
		CookieMaxAge:      time.Hour * 336,
		SecureCookie:      false,
		CookieHTTPOnly:    false,
		CookieDomain:      "neteralex.cn",
		CookieName:        "jwt-cookie",
		CookieSameSite:    protocol.CookieSameSiteDisabled,
		SendAuthorization: true,
		DisabledAbort:     false,
	})
	if err != nil {
		log.Fatalf("[JWT] Failed to init JWT middleware:%v", err)
	}
	return authMiddleware
}
