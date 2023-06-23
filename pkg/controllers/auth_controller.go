package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/hellskater/omniblog/pkg/repositories"
	"github.com/hellskater/omniblog/pkg/utils"
	"github.com/hellskater/omniblog/pkg/validators"
	"github.com/jinzhu/gorm"
)

// AuthController is a controller for authentication
type AuthController struct {
	userRepo      repositories.UserRepositoryInterface
	userValidator validators.UserValidatorInterface
}

// NewAuthController is a constructor for AuthController
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		userRepo:      repositories.NewUserRepository(db),
		userValidator: validators.NewUserValidator(),
	}
}

// Authenticate is a method for authentication
func (controller *AuthController) Authenticate(c *gin.Context) {

}

func (controller *AuthController) SignUp(c *gin.Context) {
	var user models.User
	utils.ValidateJSON(c, &user)

	err := controller.userValidator.ValidateSignUp(&user)
	if err != nil {
		utils.SendErrorMessage(
			c,
			err.Error(),
			400,
		)
		return
	}
}
