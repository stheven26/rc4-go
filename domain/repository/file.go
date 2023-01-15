package repository

import (
	"hashing-file/domain/entity"

	"gorm.io/gorm"
)

type fileRepository struct {
	DB *gorm.DB
}

type FileRepository interface {
	Upload(entity.File) (*entity.File, error)
	Encrypt(entity.File) (*entity.File, error)
	Decrypt(entity.File) (*entity.File, error)
	GetAllDocument() (int64, error)
}

func InitFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{DB: db}
}

func (f *fileRepository) Upload(data entity.File) (*entity.File, error) {
	if err := f.DB.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (f *fileRepository) Encrypt(data entity.File) (*entity.File, error) {
	if err := f.DB.Model(&data).Where(&entity.File{ID: data.ID}).Updates(data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (f *fileRepository) Decrypt(data entity.File) (*entity.File, error) {
	if err := f.DB.Model(&data).Where(&entity.File{ID: data.ID}).Updates(data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (f *fileRepository) GetAllDocument() (data int64, err error) {
	file := []entity.File{}
	if err = f.DB.Find(&file).Count(&data).Error; err != nil {
		return
	}
	return
}
