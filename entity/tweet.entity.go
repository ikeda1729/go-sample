package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	ID        int64 `gorm:"primary_key:auto_increment" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Content   string         `gorm:"type:text" json:"-"`
	UserID    int64          `gorm:"not null" json:"-"`
	User      User           `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
