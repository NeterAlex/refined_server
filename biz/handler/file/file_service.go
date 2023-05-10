package file

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
	"net/http"
)

func AvatarUpload(ctx context.Context, c *app.RequestContext) {
	file, _ := c.FormFile("file")
	userID := c.Param("id")
	log.Printf("Receiving an avatar named: %v\n", file.Filename)
	err := c.SaveUploadedFile(file, fmt.Sprintf("static/avatar/%s.jpg", userID))
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"Code": "1",
			"Msg":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, utils.H{
		"Code": "0",
		"Msg":  "File upload success",
	})
}

func AvatarDownloadHandler(ctx context.Context, c *app.RequestContext) {
	userID := c.Param("id")
	c.File(fmt.Sprintf("static/avatar/%s.jpg", userID))
}
