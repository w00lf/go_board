package main

import (
	"fmt"
	"w00lf/go_board/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "w00lf/go_board/Godeps/_workspace/src/github.com/lib/pq"
)

func inititalizeDb() gorm.DB {
  fmt.Println(ApplicationConf)
	var db, err = gorm.Open(ApplicationConf.Database["Adapter"], ApplicationConf.Database["connectString"])
	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&Post{})
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}
