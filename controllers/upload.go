package controllers

import (
	"bee-demo/models"
	"bee-demo/pkg/tencentcos"
	"bee-demo/utils"
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type UploadController struct {
	web.Controller
}

// @Title 文件上传
// @Description 文件上传
// @Failure 500
// @router / [post]
func (c *UploadController) Upload() {
	f, h, err := c.GetFile("file")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "上传失败", map[string]string{"error": err.Error()}, 500, 500)
	}
	defer f.Close()

	relFile, err := h.Open()
	if err != nil {
		models.RespondWithJSON(&c.Controller, "上传失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	filePath := "/file/" + fmt.Sprintf("%d", time.Now().Unix()) + "." + strings.Split(h.Filename, `.`)[1]

	res, err := tencentcos.UploadFile(relFile, filePath)

	if err != nil {
		models.RespondWithJSON(&c.Controller, "上传失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	url := utils.GetConfigValue("tencent_cos::COS_URL") + filePath

	data := map[string]interface{}{"url": url, "res": res.Status}

	models.RespondWithJSON(&c.Controller, "上传成功", data)
}
