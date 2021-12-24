package ik_common

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

// use a singleton of validator, it caches struct info
var instance *validator.Validate
var once sync.Once
func GetValidator() *validator.Validate {
  once.Do(func() {instance = validator.New()})
  return instance
}

type MerchCategory struct {
  Id          int    `json:"id" db:"id"`
  Title       string `json:"title" db:"title" validate:"required,max=20"`
  Description string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
}

func (m *MerchCategory) Validate() error {
  return GetValidator().Struct(*m)
}

type MerchCategoryUpdate struct {
  Title       string `json:"title,omitempty" db:"title" validate:"omitempty,max=20"`
  Description string `json:"description,omitempty" db:"description" validate:"required_without=Title,max=40"`
}

func (m *MerchCategoryUpdate) Validate() error {
  return GetValidator().Struct(*m)
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

/*type ErrorValidate struct {
  Error       string  `json:"error"`
  Field       string  `json:"field"`
  Namespace   string  `json:"namespace"`
  Param       string  `json:"param"`
  Tag         string  `json:"tag"`
}

func ValidateStruct(merchCategory interface{}) []*ErrorValidate {
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
  } else {
    ret = append(ret, &ErrorValidate{Error: "Unknown struct"})
  }

  return ret
}*/