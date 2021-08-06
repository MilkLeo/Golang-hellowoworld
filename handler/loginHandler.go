package handler

import (
	"HelloWorld/utils"
	"fmt"
	"github.com/kataras/iris/v12"
)

func UserHandle(ctx iris.Context) {

	if ctx.Path() == "/user/login" {
		ctx.Next()

	} else {

		logininfo, err := utils.CheckToken(ctx.GetHeader("Authorization"), "helloworld")
		tokens, errs := utils.GetValue("login_" + logininfo.Uid)

		if err != nil || errs != nil {
			panic("无权登录")
		}
		fmt.Print(logininfo, "=================", tokens)
		ctx.Next()
	}
}
