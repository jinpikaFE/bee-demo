package filters

import (
	"bee-demo/utils"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/adapter/logs"
)

// JwtFilter JWT过滤器
func JwtFilter(ctx *context.Context) {
	tokenString := ctx.Input.Header("Authorization")
	if tokenString == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{"msg": "缺少token", "code": 401, "data": nil}, true, true)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := utils.CheckToken(tokenString)
	if err != nil {
		logs.Error("Token 验证失败:", err)
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{"msg": "无效的token", "code": 401, "data": nil}, true, true)
		return
	}

	// 将用户信息存入上下文
	ctx.Input.SetData("userID", claims.UserID)
	ctx.Input.SetData("userName", claims.UserName)
}
