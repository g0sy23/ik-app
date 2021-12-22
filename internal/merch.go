package ik_common

import "github.com/go-playground/validator/v10"

type MerchCategory struct {
  Id          int    `json:"id" db:"id"`
  Title       string `json:"title" db:"title" validate:"required,min=3,max=30"`
  Description string `json:"description,omitempty" db:"description"`
}

type MerchItem struct {
  Id          int    `json:"id"`
  CategoryId  int    `json:"category_id"`
  Title       string `json:"title"`
  Description string `json:"description"`
  //Images			string[]
  Quantity    int    `json:"quantity"`
  Size        int    `json:"size"`
  Color       int    `json:"color"`
}

type ErrorValidate struct {
  Error       string
  Field       string
  Namespace   string
  Param       string
  Tag         string
}

func ValidateStruct(merchCategory MerchCategory) []*ErrorValidate {
  var ret []*ErrorValidate

  validate := validator.New()
  if err := validate.Struct(merchCategory); err != nil {
    for _, err := range err.(validator.ValidationErrors) {
      var i ErrorValidate
      i.Error     = err.Error()
      i.Field     = err.Field()
      i.Namespace = err.Namespace()
      i.Param     = err.Param()
      i.Tag       = err.Tag()
      ret         = append(ret, &i)
    }
  }

  return ret
}