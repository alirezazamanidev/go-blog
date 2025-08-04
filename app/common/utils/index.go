package utils

import (
	"fmt"
	"math/rand"

)

func GenerateOtp() string {
	otp := rand.Intn(999999)
	return fmt.Sprintf("%06d", otp)
}