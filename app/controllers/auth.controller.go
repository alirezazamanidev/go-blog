package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}
type SendOtpDto struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
}

func AuthNewController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

func (s *AuthController) SendOtp(c *gin.Context) {
	var sendotpDto SendOtpDto
	if err := c.ShouldBind(&sendotpDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}
	

}
