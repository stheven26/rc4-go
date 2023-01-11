package entity

import "time"

type User struct {
	ID        int64      `gorm:"column:id; primaryKey;" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Email     string     `gorm:"column:email" json:"email" gorm:"unique"`
	Password  string     `gorm:"column:password" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}
