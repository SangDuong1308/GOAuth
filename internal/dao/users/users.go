package dao

import (
	"GOAuth/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) FindByID(id uint) (*models.User, error) {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return &user, nil
}

func (u *User) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := u.db.Where("email =?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return &user, nil
}
