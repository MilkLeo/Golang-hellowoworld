package model

import "time"

type Role struct {
	Id         int       `xorm:"<-"`
	Name       string    `json:"name"`
	Isvalid    *int      `xorm:"default (1)"`
	Status     *int      `xorm:"default (1)"`
	CreateTime time.Time `xorm:"<-"`
	UpdateTime time.Time `xorm:"<-"`
}

//映射数据表
func (Role) TableName() string {
	return "t_role"
}
