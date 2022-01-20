package ik_models

type MerchItem struct {
	Id          *int    `json:"id" db:"id"`
	CategoryId  *int    `json:"category_id,omitempty" db:"category_id" validate:"omitempty,numeric,min=1"`
	Title       *string `json:"title" db:"title" validate:"required,max=20"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
	Price       *int    `json:"price" db:"price" validate:"required,numeric,min=0"`
	Quantity    *int    `json:"quantity" db:"quantity" validate:"required,numeric,min=0"`
	Size        *string `json:"size,omitempty" db:"size" validate:"omitempty,max=20"` // todo validate
	Color       *string `json:"color,omitempty" db:"color" validate:"omitempty,hexcolor|rgb|rgba|hsl|hsla,max=20"`
}

type MerchItemUpdate struct {
	CategoryId  *int    `json:"category_id,omitempty" db:"category_id" validate:"omitempty,numeric,min=1"`
	Title       *string `json:"title,omitempty" db:"title" validate:"omitempty,max=20"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty,max=40"`
	Price       *int    `json:"price,omitempty" db:"price" validate:"omitempty,numeric,min=0"`
	Quantity    *int    `json:"quantity,omitempty" db:"quantity" validate:"omitempty,numeric,min=0"`
	Size        *string `json:"size,omitempty" db:"size" validate:"omitempty,max=20"` // todo validate
	Color       *string `json:"color,omitempty" db:"color" validate:"omitempty,hexcolor|rgb|rgba|hsl|hsla,max=20"`
}
