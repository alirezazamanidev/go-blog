package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alirezazamanidev/go-blog/app/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthController struct {
	DB  *gorm.DB
	RDB *redis.Client
}

type SendOtpDto struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
}

func AuthNewController(db *gorm.DB, rdb *redis.Client) *AuthController {
	return &AuthController{
		DB:  db,
		RDB: rdb,
	}
}

func (s *AuthController) SendOtp(c *gin.Context) {
	var dto SendOtpDto
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	otp, err := s.CreateAndSaveOtp(dto.Phone, c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP sent successfully",
		"otp":     otp, // remove in production
	})
}

func (s *AuthController) CreateAndSaveOtp(phone string, c *gin.Context) (string, error) {
	key := fmt.Sprintf("otp:%s", phone)

	err := s.RDB.Get(c, key).Err()
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "OTP already sent. Please wait before requesting again.",
		})
		return "", errors.New("otp exists")
	}

	if err != redis.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Redis error",
		})
		return "", err
	}

	otp := utils.GenerateOtp()
	err = s.RDB.Set(c, key, otp, 5*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store OTP",
		})
		return "", err
	}

	return otp, nil
}
