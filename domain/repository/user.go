package repository

import (
	"hashing-file/domain/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Register(entity.User) (*entity.User, error)
	Login(string) (*entity.User, error)
	User(string) (*entity.User, error)
	GetAllUser() (int64, error)
}

func InitUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (u *userRepository) Register(data entity.User) (*entity.User, error) {
	if err := u.DB.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *userRepository) Login(email string) (data *entity.User, err error) {
	if err = u.DB.Where(`email=?`, email).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (u *userRepository) User(id string) (data *entity.User, err error) {
	if err = u.DB.Where("id=?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userRepository) GetAllUser() (data int64, err error) {
	if err = u.DB.Table("users").Count(&data).Error; err != nil {
		return
	}
	return
}
