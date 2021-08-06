package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/go-xorm/xorm"
	"log"
)

const DriverName = "mysql"

const MasterDataSourceName = "suewong:wsr951828649@tcp(rm-bp1t5ey75io2bzn1jfo.mysql.rds.aliyuncs.com:3306)/wolfkiller?charset=utf8"

var engine *xorm.Engine

func Init() *xorm.Engine {
	var err error
	engine, err = xorm.NewEngine("mysql", MasterDataSourceName)

	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}

	//ping数据库，可不要
	//if  err:=engine.Ping();err!=nil{
	//	panic(err)
	//
	//}

	//显示sql执行，便于调试分析
	engine.ShowSQL(true)
	//if err:=engine.Sync(new(model.User));err!=nil {
	//	log.Fatal("数据库同步失败",err)
	//}
	return engine
}

func Close(engine *xorm.Engine) {

	if engine != nil {
		err := engine.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
