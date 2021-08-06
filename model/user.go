package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
	"time"
)

type User struct {
	Id         int `xorm:"pk autoincr"`
	Username   string
	Phone      string
	Isvalid    int `xorm:"default 1"`
	Image      string
	Password   string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	Token      string    `xorm:"-"`
}

//映射数据表
func (User) TableName() string {
	return "t_user"
}
