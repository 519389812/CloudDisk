package models

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "../Cloud-disk.db")
	if err != nil {
		log.Println("Xorm new engine error:%v", err)
		return nil
	}
	return engine
}
