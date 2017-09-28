package models

import (
  "github.com/murillio4/stack-server/db"
  "golang.org/x/crypto/bcrypt"
  "time"
  "math/rand"
)

type User struct {
  ID          uint        `gorm:"primary_key"`

  Username    string      `json:"username" gorm:"not null;unique_index;size:255" validate:"required,min=1,max=255"`
  Password    string      `json:"password" gorm:"not null;" validate:"required,min=8"`
  Email       string      `json:"email" gorm:"not null;unique_index"`

  Confirmed   bool        `json:"-"`
  ConfirmMail ConfirmMail `json:"-"`

  RwtToken    string      `json:"-"`
}

type ConfirmMail struct {
  ID          uint        `gorm:"primary_key"`
  CreatedAt   time.Time   `json:"-"`
  UserID      uint

  Token       int         `grom:"not null;"`
}

func init() {
  db.Orm.AutoMigrate(&User{})
  db.Orm.AutoMigrate(&ConfirmMail{})
}

func SignInUser(u *User) (*Error) {
  password := u.Password

  db.Orm.First(u)

  if err := checkPassword(u, password); err != nil {
    return &Error{401, "Wrong password or username"}
  }

  return nil
}

func CreateUser(u *User) (*Error) {
  if v := Validate(u); v != nil {
    return v
  }

  if err := cryptPassword(u); err != nil {
    return err
  }

  u.ConfirmMail = ConfirmMail{Token:rand.Intn(9999)}

  db.Orm.Create(u)

  return nil
}

func AddRwtToUser(u *User, t string) (*Error) {
  if t == "" || u == nil {
    return &Error{500, "Internal server error"}
  }

  db.Orm.Model(&u).Update("RwtToken", t)

  return nil
}

func GetUserWithRwt(t string) (*User){
  u := &User{RwtToken:t}
  db.Orm.Where(&u).First(u)

  return u
}

func cryptPassword(u *User) (*Error){
  hPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return &Error{500, "Server error"}
  }

  u.Password = string(hPass[:])

  return nil
}

func checkPassword(u *User, password string) (error) {
  return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
