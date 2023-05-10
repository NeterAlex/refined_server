package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func AdminCheckService(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, utils.H{
		"Code": "0",
		"Msg":  "admin",
	})
}
