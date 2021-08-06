package dao

import (
	"HelloWorld/model"
	"log"
)
import "HelloWorld/database"

func InsertRole(role *model.Role) bool {

	engine := database.Init()
	if role == nil {
		log.Fatal("role is null")
	}
	if role.Isvalid == nil {
		isvalid := 1
		role.Isvalid = &isvalid
	}
	if role.Status == nil {
		status := 1
		role.Status = &status
	}

	count, err := engine.Insert(role)

	if count <= 0 {
		log.Fatal(err)
		return false
	}
	return true
}

func QueryAllRole() []model.Role {
	roles := make([]model.Role, 0)
	engine := database.Init()
	engine.SQL("select id,name,isvalid,status,create_time,update_time from t_role where isvalid=1").Find(&roles)
	database.Close(engine)
	return roles
}
