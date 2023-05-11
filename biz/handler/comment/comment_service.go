// Code generated by hertz generator.

package comment

import (
	"Refined_service/biz/dal/sqlite"
	"Refined_service/biz/model/comment"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UpdateComment .
// @Summary 更新评论
// @Produce json
// @Param author body string true "作者"
// @Param content body string true "内容"
// @router /v1/comment/update/:id [PUT]
func UpdateComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.UpdateCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusOK, &comment.UpdateCommentResponse{Code: comment.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	cm := &comment.Comment{}
	cm.Author = req.Author
	cm.Content = req.Content
	cm.ID = req.ID
	cm.UserID = req.UserID
	cm.Date = req.Date

	if err = sqlite.Update[comment.Comment](cm.ID, cm); err != nil {
		c.JSON(http.StatusOK, &comment.UpdateCommentResponse{Code: comment.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &comment.UpdateCommentResponse{Code: comment.Code_Success})
}

// DeleteComment .
// @Summary 删除评论
// @Produce json
// @Param id path string true "评论id"
// @router /v1/comment/delete/:id [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.DeleteCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &comment.DeleteCommentResponse{Code: comment.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = sqlite.Delete[comment.Comment](req.ID); err != nil {
		c.JSON(consts.StatusOK, &comment.DeleteCommentResponse{Code: comment.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(consts.StatusOK, comment.DeleteCommentResponse{Code: comment.Code_Success})
}

// QueryComment .
// @Summary 查询评论
// @Produce json
// @Param id body string true "文章ID"
// @Param page body string true "页码"
// @Param page_size body string true "页面大小"
// @router /v1/comment/query/ [GET]
func QueryComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.QueryCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &comment.QueryCommentResponse{Code: comment.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	comments, total, err := sqlite.Query[comment.Comment]("post_id = ?", *req.ID)
	if err != nil {
		c.JSON(consts.StatusOK, &comment.QueryCommentResponse{Code: comment.Code_DbError, Msg: err.Error()})
		return
	}
	c.JSON(consts.StatusOK, &comment.QueryCommentResponse{Code: comment.Code_Success, Comments: comments, Total: total})
}

// CreateComment
// @Summary 创建评论
// @Produce json
// @Param author body string true "作者"
// @Param content body string true "内容"
// @Param postID body string true "所属文章id"
// @router /v1/comment/create/ [POST]
func CreateComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.CreateCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &comment.CreateCommentResponse{Code: comment.Code_ParamInvalid, Msg: err.Error()})
		return
	}
	if err = sqlite.Create[comment.Comment]([]*comment.Comment{
		{
			Author:  req.Author,
			Content: req.Content,
			PostID:  req.PostID,
			UserID:  req.UserID,
			Date:    req.Date,
		},
	}); err != nil {
		c.JSON(consts.StatusOK, &comment.CreateCommentResponse{Code: comment.Code_DbError, Msg: err.Error()})
		return
	}

	c.JSON(consts.StatusOK, comment.CreateCommentResponse{Code: comment.Code_Success})
}
