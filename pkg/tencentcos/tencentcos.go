package tencentcos

import (
	"bee-demo/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"

	"context"
	"io"
)

var (
	Client *cos.Client
)

func init() {
	u, _ := url.Parse(utils.GetConfigValue("tencent_cos::COS_URL"))
	log.Println((utils.GetConfigValue("tencent_cos::COS_URL")))
	b := &cos.BaseURL{BucketURL: u}
	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  fmt.Sprintf("%s", utils.GetConfigValue("tencent_cos::SECRET_ID")),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: fmt.Sprintf("%s", utils.GetConfigValue("tencent_cos::SECRET_KEY")), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
}

func UploadFile(file io.Reader, filePath string) (*cos.Response, error) {
	res, err := Client.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {

		return nil, err
	}
	return res, nil
}
