package ik_models

import (
	"errors"
	"sync"

	"github.com/go-playground/validator/v10"
)

func (m *MerchCategory) Validate() error {
	return GetValidator().Struct(*m)
}

func (m *MerchCategoryUpdate) Validate() error {
	if *m == (MerchCategoryUpdate{}) {
		return errors.New("empty content")
	}
	return GetValidator().Struct(*m)
}

func (m *MerchItem) Validate() error {
	return GetValidator().Struct(*m)
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
	once.Do(func() { instance = validator.New() })
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