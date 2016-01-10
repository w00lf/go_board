package main

import (
  "github.com/jinzhu/gorm"
  "fmt"
   _ "github.com/lib/pq"
)

func inititalizeDb() gorm.DB {
  var db, err = gorm.Open("postgres", "dbname=go_board user=mik password=123 host=localhost sslmode=disable")  
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