package main

import (
	"CloudDisk/models"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	fmt.Println("Start.")
	var err error
	engine, err = xorm.NewEngine("sqlite3", "Cloud-disk.db")
	if err != nil {
		fmt.Println("engine error")
		fmt.Println(err)
	}
	data := make([]*models.UserBasic, 0)

	//tbs, err := engine.DBMetas()
	//if err != nil {
	//	fmt.Println("db error")
	//	fmt.Println(err)
	//	return
	//}
	//for i, tb := range tbs {
	//	fmt.Println("index:", i, "tbName", tb.Name)
	//}
	err = engine.Find(&data)
	if err != nil {
		fmt.Println("find error")
		fmt.Println(err)
	}
	fmt.Println(data)

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal error")
		fmt.Println(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		fmt.Println("indent error")
		fmt.Println(err)
	}
	fmt.Println(dst.String())
}
