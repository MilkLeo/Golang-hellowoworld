package router

import "github.com/kataras/iris/v12"
import "HelloWorld/controller"

func RegisterUserRouter(app *iris.Application) {
	app.Get("/user/queryUser", controller.QueryUserInfo)
	app.Get("/user/QueryUserInfoById", controller.QueryUserInfoById)
	app.Get("/user/login", controller.Login)

}
