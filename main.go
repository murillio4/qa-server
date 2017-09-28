package main

import (
	//"fmt"
	"github.com/murillio4/stack-server/db"
	_ "github.com/murillio4/stack-server/routers"
	//"github.com/murillio4/stack-server/models"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	defer db.Orm.Close()
	//question := models.Question{Title:"asjdnjasdnkjasnddkaj", Content:"fittmasdjknaskjdnkasjndkajsasdasdasdasdsadsadasdasdasdasdasdndkjasndkajdnkj"}
	// pagination := models.Pagination{PerPage:3, CurrPage:3}
	//comment := models.Comment{Content:"sss2323232323sss"}
	//answer := models.Answer{Content:"lkmkmsdsdflkmmsdlfkmsjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjnjsndkasndknasdknaskdjnaskjdnaskjdnaskdjnaskjdndskjn"}
	//db.Orm.Create(&question)
	//db.Orm.First(&question, 1).Create(&answer)
	//db.Orm.First(&question, 1)
	//db.Orm.First(&question, 1).Association("Comment").Append(&comment)
	//db.Orm.Create(&answer)
	//db.Orm.Preload("Answer").First(&question, 1)
	//db.Orm.First(&models.Answer{}, 4).Association("Comment").Append(&comment)
	//db.Orm.Preload("Answer.Comment").Preload("Answer").Preload("Comment").First(&question)
	//models.CreateQuestion(&question)
	//fmt.Println(models.)
	// models.GetAnswers(question, &pagination)
	//models.CreateAnswer(question, &answer)
	// fmt.Printf("%+v\n", pagination)
	//fmt.Println(db.Orm.NewRecord(question))
	beego.Run()
}
