package models

import (
  "github.com/murillio4/stack-server/db"
  "time"
)

type Question struct {
  ID        uint        `json:"id" gorm:"primary_key"`
  CreatedAt time.Time   `json:"created_at"`
  UpdatedAt time.Time   `json:"updated_at"`
  DeletedAt *time.Time  `json:"deleted_at"`

  Title     string      `json:"title" gorm:"not null;size:255" validate:"required,min=10,max=255"`
  Content   string      `json:"content" gorm:"not null" validate:"required,min=30"`
  Viewed    int         `json:"viewed" gorm:"default:0"`

  //has multiple answers and comments
  Answer    []Answer    `json:"answers"`
  Comment   []Comment   `json:"comments" gorm:"polymorphic:Commentable;"`
}


func init() {
  db.Orm.AutoMigrate(&Question{})
}

func GetQuestion(q *Question) (*Error) {
  db.Orm.Preload("Comment").First(q)

  if q.Title == "" && q.Content == "" {
    return &Error{404, "Could not find question"}
  }

  return nil
}

func RegisterView(q Question) {
  //Do this in GetQuestion??
  db.Orm.First(&q).UpdateColumn("viewed", q.Viewed)
}

func CreateQuestion(q *Question) (*Error){
  if v := Validate(q); v != nil {
    return v
  }

  db.Orm.Create(&q)
  return nil
}

func UpdateQuestion(q *Question) (*Error){
  if v := Validate(q); v != nil {
    return v
  }

  temp := q
  db.Orm.First(&temp)

  temp.Title = q.Title
  temp.Content = q.Content

  db.Orm.Save(q)
  return nil
}

func DeleteQuestion(q Question) {

}
