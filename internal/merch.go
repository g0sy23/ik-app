package ik_common

import (
  "errors"
  "sync"

  "github.com/go-playground/validator/v10"
)

type MerchCategory struct {
  Id          *int    `json:"id" db:"id"`
  Title       *string `json:"title" db:"title" validate:"required,max=20"`
  Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
}

func (m *MerchCategory) Validate() error {
  return GetValidator().Struct(*m)
}

type MerchCategoryUpdate struct {
  Title       *string `json:"title,omitempty" db:"title" validate:"omitempty,max=20"`
  Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
}

func (m *MerchCategoryUpdate) Validate() error {
  if *m == (MerchCategoryUpdate{}) {
    return errors.New("empty content")
  }
  return GetValidator().Struct(*m)
}

type MerchItem struct {
  Id          *int    `json:"id" db:"id"`
  CategoryId  *int    `json:"category_id,omitempty" db:"category_id" validate:"omitempty,numeric,min=1"`
  Title       *string `json:"title" db:"title" validate:"required,max=20"`
  Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
  //Images			string[]
  Price       *int    `json:"price" db:"price" validate:"required,numeric,min=0"`
  Quantity    *int    `json:"quantity" db:"quantity" validate:"required,numeric,min=0"`
  Size        *string `json:"size,omitempty" db:"size" validate:"omitempty,max=20"` // todo validate
  Color       *string `json:"color,omitempty" db:"color" validate:"omitempty,hexcolor|rgb|rgba|hsl|hsla,max=20"`
}

func (m *MerchItem) Validate() error {
  return GetValidator().Struct(*m)
}

type MerchItemUpdate struct {
  CategoryId  *int    `json:"category_id,omitempty" db:"category_id" validate:"omitempty,numeric,min=1"`
  Title       *string `json:"title,omitempty" db:"title" validate:"omitempty,max=20"`
  Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
  //Images			string[]
  Price       *int    `json:"price,omitempty" db:"price" validate:"omitempty,numeric,min=0"`
  Quantity    *int    `json:"quantity,omitempty" db:"quantity" validate:"omitempty,numeric,min=0"`
  Size        *string `json:"size,omitempty" db:"size" validate:"omitempty,max=20"` // todo validate
  Color       *string `json:"color,omitempty" db:"color" validate:"omitempty,hexcolor|rgb|rgba|hsl|hsla,max=20"`
}

func (m *MerchItemUpdate) Validate() error {
  if *m == (MerchItemUpdate{}) {
    return errors.New("empty content")
  }
  return GetValidator().Struct(*m)
}

// use a singleton of validator, it caches struct info
var instance *validator.Validate
var once sync.Once
func GetValidator() *validator.Validate {
  once.Do(func() {instance = validator.New()})
  return instance
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