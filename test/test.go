package main

import (
	"CloudDisk/models"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "/Cloud-disk.db")
	if err != nil {
		return
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		return
	}
	fmt.Println(data)
}
