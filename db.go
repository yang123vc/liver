package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDataBase() {
	tmplString := "{{.User}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.Database}}?charset=utf8mb4&parseTime=True&loc=Local"
	database, _ := tmplToString(tmplString, cfg.DB)
	var err error
	db, err = gorm.Open("mysql", database)
	if err != nil {
		panic(err.Error())
	}
}
