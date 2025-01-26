package handler

import (
	"log"
	"myapp/db"
	"myapp/helper"
	"myapp/model"
	"myapp/repo"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register handler untuk proses pendaftaran user baru
func Register(c echo.Context) error {
	var user model.User

	// Binding request body ke struct user
	if err := c.Bind(&user); err != nil {
		log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, helper.Error("Invalid request body", http.StatusBadRequest))
	}

	// Hash password sebelum menyimpan
	hashedPassword, err := helper.HashedPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Error("Failed to hash password", http.StatusInternalServerError))
	}

	// Update password yang sudah di-hash
	user.Password = hashedPassword

	// Menggunakan repo untuk menyimpan user baru
	userRepo := repo.NewUserRepo(db.DB) // Pastikan db.DB sudah terhubung
	err = userRepo.Register(user)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return c.JSON(http.StatusInternalServerError, helper.Error("Failed to register user", http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, helper.Succes("User registered successfully", http.StatusOK, user))
}

// Login handler untuk proses login
func Login(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Binding request body ke loginRequest struct
	if err := c.Bind(&loginRequest); err != nil {
	log.Printf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, helper.Error("Invalid request body", http.StatusBadRequest))
	}

	// Menggunakan repo untuk mendapatkan user berdasarkan email
	userRepo := repo.NewUserRepo(db.DB) // Pastikan db.DB sudah terhubung
	user, err := userRepo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return c.JSON(http.StatusInternalServerError, helper.Error("Failed to get user", http.StatusInternalServerError))
	}

	// Verifikasi password dengan hash
	if !helper.CheckPasswordHash(loginRequest.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, helper.Error("Invalid email or password", http.StatusUnauthorized))
	}

	// Generate JWT token
	token, err := helper.GenerateJWT(user.Id, user.Email)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return c.JSON(http.StatusInternalServerError, helper.Error("Failed to generate JWT", http.StatusInternalServerError))
	}

	// Response dengan token dan data user
	response := map[string]interface{}{
		"token":   token,
		"status" : http.StatusOK,
		"message": "Login successful",
		"user": map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	return c.JSON(http.StatusOK, response)
}