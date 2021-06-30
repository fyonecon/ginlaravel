package Middlewares
// 处理web_token
// 通常情况下：1. SEO页面都是公开数据，没必要用CSRF；2. 登录与用户数据页面都是前后端分离。所以"模版+模版登录"是一个淘汰的功能，"模版+Api登录"可能会有晚霞余光。

import "github.com/gin-gonic/gin"

func CreateWebToken(ctx *gin.Context) string {
	return "csrf"
}

func VerifyWebToken(ctx *gin.Context)  {
	ctx.Next()
}
