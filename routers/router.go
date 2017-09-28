// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	c "github.com/murillio4/stack-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/questions",
			beego.NSRouter("/:qid([0-9]+)", &c.QuestionController{}, "get:GetQuestion"),
			beego.NSRouter("/", &c.QuestionController{}, "post:CreateQuestion"),
		),
		beego.NSNamespace("/answers",
			beego.NSRouter("/:qid([0-9]+)/paginate", &c.AnswerController{}, "get:GetPaginateAnswers"),
			beego.NSRouter("/:aid([0-9]+)", &c.AnswerController{}, "get:GetAnswer"),
			beego.NSRouter("/:qid([0-9]+", &c.AnswerController{}, "post:CreateAnswer"),
		),
		beego.NSNamespace("/comments",
			beego.NSRouter("/answer/:aid([0-9]+)", &c.CommentController{}, "post:CreateAnswerComment"),
			beego.NSRouter("/question/:qid([0-9]+)", &c.CommentController{}, "post:CreateQuestionComment"),
		),
		beego.NSNamespace("/users",
			beego.NSRouter("/register", &c.UserController{}, "post:Register"),
			beego.NSRouter("/login", &c.UserController{}, "post:SignIn"),
		),
	)
	beego.AddNamespace(ns)
}
