package repo

import (
	"fmt"
	"myapp/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Register(user model.User) error
	Login(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo: Constructor untuk UserRepo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

// Register: Menyimpan data user baru
func (r *userRepo) Register(user model.User) error {
	return r.db.Create(&user).Error
}

// Login: Memeriksa apakah user ada berdasarkan email dan password
func (r *userRepo) Login(user model.User) (model.User, error) {
	// Mengambil user berdasarkan email
	existingUser, err := r.GetUserByEmail(user.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("email atau password salah: %w", err)
	}

	// Mengecek password (anda bisa melakukan hashing disini jika diperlukan)
	if existingUser.Password != user.Password {
		return model.User{}, fmt.Errorf("email atau password salah")
	}

	return existingUser, nil
}

func (r *userRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
