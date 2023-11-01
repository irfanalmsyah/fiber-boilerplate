package model

import (
    "gorm.io/gorm"
    "github.com/go-playground/validator/v10"
)

type User struct {
    gorm.Model           // Adds some metadata fields to the table
    ID          uint      `gorm:"primaryKey;autoIncrement"`
    Name        string    `gorm:"size:255;not null;"`
    Email       string    `gorm:"size:255;not null;uniqueIndex"`
    Password    string    `gorm:"size:255;not null;"`
}

type SignUpInput struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    ConfirmPassword string `json:"confirmPassword" validate:"required,min=6"`
}

type LoginInput struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
