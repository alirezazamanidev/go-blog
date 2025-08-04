package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateOtp() string {
	otp := rand.Intn(999999)
	return fmt.Sprintf("%06d", otp)
}

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwt(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * 7 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
