package models

import (
  "github.com/murillio4/stack-server/db"
  "time"
)

type Comment struct {
  ID        uint            `json:"id" gorm:"primary_key"`
  CreatedAt time.Time       `json:"created_at"`
  UpdatedAt time.Time       `json:"updated_at"`
  DeletedAt *time.Time      `json:"deleted_at"`

  //belongs to
  CommentableId   int       `json:"-"`
  CommentableType string    `json:"-"`

  Content         string    `json:"content" gorm:"not null;" validate:"required,min=5,max=300"`
}

func init() {
  db.Orm.AutoMigrate(&Comment{})
}

func CreateComment(p interface{}, c *Comment) (*Error){
  if v := Validate(c); v != nil {
    return v
  }

  db.Orm.Model(p).Association("Comment").Append(c)

  return nil
}
