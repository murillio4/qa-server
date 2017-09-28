package events

import (
  "github.com/murillio4/stack-server/db"
  m "github.com/murillio4/stack-server/models"
  "github.com/sadlil/go-trigger"
  "gopkg.in/gomail.v2"
  "strconv"
)

func init() {
  trigger.On("send-confirm-email", sendConfirmEmail)
  trigger.On("register-view", registerView)
}

func sendConfirmEmail(u *m.User) {
  m := gomail.NewMessage()
  m.SetHeader("From", "noreply@prognett.org")
  m.SetHeader("To", u.Username + "@uio.no")
  m.SetHeader("Subject", "Prognett Stack Confirmation")
  m.SetBody("text/html", "<h1>" + strconv.Itoa(u.ConfirmMail.Token) + "</h1>")

  d := gomail.NewDialer("mail.prognett.org", 587, "noreply@prognett.org", "Jagerte12")

  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }
}

//Can make this better, but how?? stateless api fak
func registerView(q *m.Question) {
  db.Orm.Model(q).Update("Viewed", q.Viewed+1)
}
