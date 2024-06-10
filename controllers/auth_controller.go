package controllers

import (
	"dbo/models"
	"dbo/services"
	"net/http"

	i "dbo/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

// swagger:response tokenResponse
type tokenResponseWrapper struct {
	// Example: {"token": "your-token-value"}
	Token string `json:"token"`
}

var _ = tokenResponseWrapper{}
var _ = i.ErrorResponseWrapper{}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 500 {object} i.ErrorResponseWrapper
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := ctrl.service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

// Login godoc
// @Summary Login a user
// @Description Login a user and get a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user body models.User true "User"
// @Success 200 {object} tokenResponseWrapper
// @Failure 400 {object} i.ErrorResponseWrapper
// @Failure 401 {object} i.ErrorResponseWrapper
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
