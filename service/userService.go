package service

import (
	"HelloWorld/dao"
	_ "HelloWorld/database"
	"HelloWorld/model"
	"HelloWorld/utils"
	"fmt"
	"strconv"
	"time"
)

//添加用户
func insert() {

}

//查询所有用户
func QueryAllUser() []model.User {

	return dao.QueryAllUser()
}

//查询单个用户信息
func QueryUserById(userId int) *model.User {

	return dao.QueryUserById(userId)
}

//根据用户手机号查询用户信息
func Login(phone string, password string) *model.User {

	if phone == "" || password == "" {
		return nil
	}

	user := dao.QueryUserInfoByPhone(phone)

	if user == nil {
		return nil
	}

	fmt.Print("===========" + user.Password + "================" + utils.MD5(password))
	if user.Password == utils.MD5(password) {

		userinfo := utils.LoginInfo{strconv.Itoa(user.Id), 1}
		expireTime := time.Now().Add(time.Hour * 24).Unix()

		var err error
		user.Token, err = utils.GenerateToken(userinfo.Uid, userinfo.Utype, "helloworld", expireTime)

		flag := utils.SetValue("login_"+strconv.Itoa(user.Id), user.Token, 3*time.Hour)
		if err != nil || !flag {
			fmt.Print(err)
		}
		return user
	}

	return nil
}
