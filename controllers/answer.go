package controllers

import (
		m "github.com/murillio4/stack-server/models"
)

type AnswerController struct {
	BaseController
}

func (c *AnswerController) GetAnswer() {

}

func (c *AnswerController) GetPaginateAnswers() {
	qid, _ := c.GetInt(":qid")

	limit, le := c.GetInt("limit")
	page, pe := c.GetInt("page")

	if le != nil {
		limit = 10
	}

	if pe != nil {
		page = 1
	}

	p := m.Pagination{PerPage:limit, CurrPage:page}
	e := m.GetPaginateAnswers(qid, &p)

	if e != nil {
    c.Data["json"] = e

  } else {
		if p.NextPage != "" {
			p.NextPage = c.URLFor("AnswerController.GetPaginateAnswers", ":qid", qid, "limit", p.PerPage, "page", p.NextPage)
		}
		if p.PrevPage != "" {
			p.PrevPage = c.URLFor("AnswerController.GetPaginateAnswers", ":qid", qid, "limit", p.PerPage, "page", p.PrevPage)
		}

    c.Data["json"] = p
  }
  c.ServeJSON()
}

func (c *AnswerController) CreateAnswer() {
	qid, _ := c.GetInt(":qid")
	content := c.GetString("a_content")

	a := m.Answer{Content:content}
	e := m.CreateAnswer(qid, &a)

	if e != nil {
		c.Data["json"] = e
	} else {
		c.Data["json"] = a
	}
	c.ServeJSON()
}
