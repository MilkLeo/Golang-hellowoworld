package dao

import (
	"HelloWorld/database"
	"HelloWorld/model"
	"fmt"
)

//查询所有用户
func QueryAllUser() []model.User {

	users := make([]model.User, 0)
	engine := database.Init()
	engine.SQL("select id,username,phone,isvalid,image,create_time,update_time from t_user where isvalid=1").Find(&users)
	fmt.Print(users)
	database.Close(engine)
	return users
}

//根据主键id查询用户信息
func QueryUserById(userId int) *model.User {

	user := new(model.User)
	engine := database.Init()
	engine.Id(userId).Get(user)
	database.Close(engine)
	return user
}

//根据用户手机号查询用户信息
func QueryUserInfoByPhone(phone string) *model.User {

	if phone == "" {
		return nil
	}

	user := new(model.User)
	engine := database.Init()
	engine.Where("phone=?", phone).Get(user)
	database.Close(engine)
	return user

}
