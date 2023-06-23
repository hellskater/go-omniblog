package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/hellskater/omniblog/pkg/models"
)

type UserValidatorInterface interface {
	ValidateSignUp(user *models.User) error
}

type UserValidator struct {
	validator *validator.Validate
}

func NewUserValidator() *UserValidator {
	return &UserValidator{validator: validator.New()}
}

func (v *UserValidator) ValidateSignUp(user *models.User) error {
	return v.validator.Struct(user)
}
