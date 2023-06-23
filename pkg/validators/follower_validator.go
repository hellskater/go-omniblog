package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/hellskater/omniblog/pkg/models"
)

type FollowerValidatorInterface interface {
	ValidateCreate(follower *models.Follower) error
}

type FollowerValidator struct {
	validator *validator.Validate
}

func NewFollowerValidator() *FollowerValidator {
	return &FollowerValidator{validator: validator.New()}
}

func (v *FollowerValidator) ValidateCreate(follower *models.Follower) error {
	return v.validator.Struct(follower)
}
