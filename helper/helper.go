package helper

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)



type FormatSuccess struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    interface{} `json:"data"`
}


type FormatError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}



func Succes(message string, status int, data interface{}) FormatSuccess {
	res := FormatSuccess{
		Status: status,
		Message:  message,
		Data: data,
	}

	return res
}



func Error(message string, status int) FormatError {
	res := FormatError{
		Status: status,
		Message:  message,
	}

	return res
}


func HashedPassword(password string) (string, error) {
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}



func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var SecretKey = os.Getenv("JWT_SECRET_KEY")

// GenerateJWT menghasilkan JWT token
func GenerateJWT(userID int, email string) (string, error) {
	// Membuat klaim (claims) untuk JWT
	claims := jwt.MapClaims{
		"sub":  userID,
		"email": email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	}

	// Membuat token JWT dengan klaim dan secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}