package entity

import "time"

type File struct {
	ID            int64      `gorm:"column:id; primaryKey;" json:"id"`
	UploadedFile  string     `gorm:"column:uplodedFile" json:"uploadedFile"`
	EncryptedFile string     `gorm:"column:encryptedFile" json:"encryptedFile"`
	DecryptedFile string     `gorm:"column:decryptedFile" json:"decryptedFile"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}
