package db

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
  Orm *gorm.DB
)

func init() {
  var err error

  Orm, err = gorm.Open("mysql", "app:app@/stack?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    log.Fatal(err)
  }
}
