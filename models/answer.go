package models

import (
  "github.com/murillio4/stack-server/db"
  "time"
)

type Answer struct {
  ID        uint        `json:"id" gorm:"primary_key"`
  CreatedAt time.Time   `json:"created_at"`
  UpdatedAt time.Time   `json:"updated_at"`
  DeletedAt *time.Time  `json:"deleted_at"`

  //belongs to question
  QuestionID  int       `json:"-"`

  Content     string    `json:"content" gorm:"not null;" validate:"required,min=30"`

  //has many comments
  Comment     []Comment `json:"comments,omitempty" gorm:"polymorphic:Commentable;"`
}

func init() {
  db.Orm.AutoMigrate(&Answer{})
}

func GetPaginateAnswers(qid int, p *Pagination) (*Error){
  q := Question{ID:uint(qid)}
  p.Total = db.Orm.Model(&q).Association("Answer").Count()
  e := paginate(p)

  if e != nil {
    return e
  }

  p.Content = &[]Answer{}
  db.Orm.Model(&q).Offset(p.offset).Limit(p.PerPage).Preload("Comment").Association("Answer").Find(p.Content)

  return nil
}


func CreateAnswer(qid int, a *Answer) (*Error){
  if v := Validate(a); v != nil {
    return v
  }

  q := Question{ID:uint(qid)}
  db.Orm.Model(&q).Association("Answer").Append(a)

  return nil
}
