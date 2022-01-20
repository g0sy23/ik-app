package ik_models

type MerchCategory struct {
	Id          *int    `json:"id" db:"id"`
	Title       *string `json:"title" db:"title" validate:"required,max=20"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
}

type MerchCategoryUpdate struct {
	Title       *string `json:"title,omitempty" db:"title" validate:"omitempty,max=20"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
}
