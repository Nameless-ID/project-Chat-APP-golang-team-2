package repository

import (
	"errors"
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    GetUserByID(id string) (*models.User, error)
    GetAllUsers() ([]models.User, error)
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db}
}

func (r *userRepo) GetUserByID(id string) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
    return &user, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := r.db.Model(models.User{}).Where("is_active = ?", true).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("users not found")
	}
    return users, nil
}
