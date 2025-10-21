package auth

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Mobile string `json:"mobile"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	var body map[string]interface{}

	var validMobile bool = false
	var validOtp bool = false

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	mobileNumber := body["mobile"]
	otpSent := body["otp"]

	re := regexp.MustCompile(`^[6-9]\d{9}$`)

	if re.MatchString((mobileNumber.(string))) {
		validMobile = true
		c.JSON(http.StatusOK, gin.H{})
		fmt.Println(mobileNumber.(string), "is a valid mobile number")
	} else {
		validMobile = false
		fmt.Println(mobileNumber.(string), "is NOT a valid mobile number")
	}

	re = regexp.MustCompile(`^\d{6}$`)

	if re.MatchString((otpSent.(string))) {

		validOtp = true
		fmt.Println(otpSent.(string), "is a valid OTP")
	} else {
		validOtp = false
		fmt.Println(otpSent.(string), "is NOT a valid OTP")
	}

	if validMobile && validOtp {

		token, _ := GenerateJWT(mobileNumber.(string))

		c.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"token":  token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "Invalid OTP",
		})
	}
}
