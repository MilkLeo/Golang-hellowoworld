package controller

import (
	"HelloWorld/model"
	"HelloWorld/service"
	_ "HelloWorld/service"
	"github.com/kataras/iris/v12"
	"log"
)

// @Summary 获取所有用户
// @Produce  json
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /user/queryUser [get]
//@Tags 用户接口
func QueryUserInfo(ctx iris.Context) {
	data := model.SuccessData(service.QueryAllUser())
	ctx.JSON(data)
}

// @Summary 获取个人用户信息
// @Produce  json
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /user/QueryUserInfoById [get]
//@Tags 用户接口
func QueryUserInfoById(ctx iris.Context) {
	userId, _ := ctx.URLParamInt("userId")
	log.Print(userId)
	data := model.SuccessData(service.QueryUserById(userId))
	ctx.JSON(data)
}

// @Summary 用户登陆
// @Produce  json
// @Success 200 {string} json "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /user/login [get]
//@Tags 用户接口
func Login(ctx iris.Context) {
	phone := ctx.URLParam("phone")
	password := ctx.URLParam("password")
	data := model.SuccessData(service.Login(phone, password))
	ctx.JSON(data)
}
