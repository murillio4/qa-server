package models

import (
  "gopkg.in/go-playground/validator.v9"
  "math"
  "strconv"
)

type Error struct {
  Code int
  Message interface{}
}

type ValidationError struct {
  Field string  `json:"field"`
  Tag   string  `json:"tag"`
}

type Pagination struct {
  Total       int         `json:"total"`
  offset      int         `json:"-"`
  TotalPages  int         `json:"total_pages"`
  PerPage     int         `json:"per_page"`


  CurrPage    int         `json:"current_page"`
  PrevPage    string      `json:"prev_page"`
  NextPage    string      `json:"next_page"`

  Content     interface{} `json:"content"`
}


func paginate(p *Pagination) (*Error) {
  p.TotalPages = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))

  if p.CurrPage < 1 || p.CurrPage > p.TotalPages {
    return &Error{404, "Page out of bounce"}
  }

  if p.CurrPage-1 >= 1 {
    p.PrevPage = strconv.Itoa(p.CurrPage-1)
  }

  if p.CurrPage+1 <= p.TotalPages {
    p.NextPage = strconv.Itoa(p.CurrPage+1)
  }

  p.offset = p.PerPage*(p.CurrPage - 1)

  return nil
}

func getErrors(err error) ([]ValidationError){
  e := []ValidationError{}

  for _, err := range err.(validator.ValidationErrors) {
    e = append(e, ValidationError{err.Field(), err.Tag()})
  }

  return e
}

func Validate(s interface{}) (*Error) {
  validate := validator.New()
  err := validate.Struct(s)

  if err != nil {
    return &Error{422, getErrors(err)}
  }

  return nil
}
