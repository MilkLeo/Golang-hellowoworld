package router

import (
	"HelloWorld/controller"
	"github.com/kataras/iris/v12"
)

func RegisterRouterRouter(app *iris.Application) {
	app.Post("/role/insertRole", controller.InsertRole)
	app.Get("/role/queryAllRole", controller.QueryAllRole)

}
