package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
)

type AuthClaim struct {
	Username string `json:"username"`
	ID       string `json:"userID"`
	Class    string `json:"status"`
	jwt.StandardClaims
}

func AdminPermissionCheck() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			token := c.Request.Header.Get("Authorization")[7:]
			//fmt.Println(token)
			parsedToken, _ := jwt.ParseWithClaims(token, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("verification"), nil
			})
			parsed := parsedToken.Claims.(*AuthClaim)
			if parsed.Class == "admin" {
				c.Next(ctx)
			} else {
				c.AbortWithMsg("Unauthorized", 401)
			}
		}}
}
