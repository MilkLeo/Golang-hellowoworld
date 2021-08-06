package controller

import (
	"HelloWorld/model"
	"HelloWorld/service"
	"github.com/kataras/iris/v12"
	"log"
)

// @Summary 添加角色
// @Produce  json
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /role/insertRole [get]
//@Tags 角色接口
func InsertRole(ctx iris.Context) {

	//2.Json数据解析
	role := &model.Role{} //通过context.ReadJSON()读取传过来的数据

	err := ctx.ReadJSON(role)

	if err != nil {
		log.Fatal(err.Error())
	}
	data := model.SuccessData(service.InsertRole(role))
	ctx.JSON(data)
}

func QueryAllRole(ctx iris.Context) {

	data := model.SuccessData(service.QueryAllRole())
	ctx.JSON(data)
}
