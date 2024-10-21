package filters

import (
	"bee-demo/utils"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
)

// JwtFilter JWT过滤器
func JwtFilter(ctx *context.Context) {
	tokenString := ctx.Input.Header("Authorization")
	if tokenString == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("缺少token"))
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := utils.CheckToken(tokenString)
	if err != nil {
		logs.Error("Token 验证失败:", err)
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("无效的token"))
		return
	}

	// 将用户信息存入上下文
	ctx.Input.SetData("userID", claims.UserID)
	ctx.Input.SetData("username", claims.Username)
}
