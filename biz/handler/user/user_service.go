// Code generated by hertz generator.

package user

import (
	"Refined_service/biz/dal/sql"
	"Refined_service/biz/pack"
	"context"
	"net/http"

	user "Refined_service/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UpdateUser .
// @Summary 更新用户
// @Produce json
// @Param id query string false "用户ID(可选)"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param nickname body string true "昵称"
// @Param email body string true "Email"
// @Param phone body string false "手机号"
// @router /v1/user/update/:id [PUT]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, &user.UpdateUserResponse{Code: user.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	u := &user.User{}
	u.ID = req.ID
	u.Username = req.Username
	u.Nickname = req.Nickname
	u.Email = req.Email
	u.Phone = req.Phone
	if req.Password != nil && *req.Password != "" {
		u.Password = pack.HashSHA256(*req.Password)
	}
	if err = sql.Update[user.User](u.ID, u); err != nil {
		c.JSON(http.StatusOK, &user.UpdateUserResponse{Code: user.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &user.UpdateUserResponse{Code: user.Code_Success})
}

// DeleteUser .
// @Summary 删除用户
// @Produce json
// @Param id path string true "用户ID"
// @router /v1/user/delete/:id [DELETE]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.DeleteUserResponse{Code: user.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = sql.Delete[user.User](req.ID); err != nil {
		c.JSON(consts.StatusOK, &user.DeleteUserResponse{Code: user.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(consts.StatusOK, user.DeleteUserResponse{Code: user.Code_Success})
}

// QueryUser .
// @Summary 查询用户
// @Produce json
// @Param id query string false "用户ID(可选)"
// @Param page query string true "页面"
// @Param page_size query string true "页面容量"
// @router /v1/user/query/ [GET]
func QueryUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.QueryUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.QueryUserResponse{Code: user.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	var users []*user.User
	var total int64
	if req.ID == "0" {
		users, total, err = sql.QueryAllExclude[user.User]("password", req.Page, req.GetPageSize())
	} else {
		users, total, err = sql.QueryExclude[user.User]("id = ?", req.ID, "password")
	}

	if err != nil {
		c.JSON(consts.StatusOK, &user.QueryUserResponse{Code: user.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(consts.StatusOK, &user.QueryUserResponse{Code: user.Code_Success, Users: users, Total: total})
}

// CreateUser .
// @Summary 创建用户
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param nickname body string true "昵称"
// @Param email body string true "Email"
// @Param phone body string false "手机号"
// @router /v1/user/create/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.CreateUserResponse{Code: user.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = sql.Create[user.User]([]*user.User{
		{
			Username: req.Username,
			Password: pack.HashSHA256(req.Password),
			Status:   "normal",
			Email:    req.Email,
			Phone:    req.Phone,
			Nickname: req.Nickname,
		},
	}); err != nil {
		c.JSON(consts.StatusOK, &user.CreateUserResponse{Code: user.Code_DbError, Msg: err.Error()})
		return
	}

	c.JSON(consts.StatusOK, user.CreateUserResponse{Code: user.Code_Success})
}
