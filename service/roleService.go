package service

import (
	"HelloWorld/dao"
	"HelloWorld/model"
)

//添加角色
func InsertRole(role *model.Role) bool {
	return dao.InsertRole(role)
}

func QueryAllRole() []model.Role {
	return dao.QueryAllRole()
}
