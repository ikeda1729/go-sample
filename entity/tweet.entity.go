package entity

type Tweet struct {
	ID      int64  `gorm:"primary_key:auto_increment" json:"-"`
	Content string `gorm:"type:text" json:"-"`
	UserID  int64  `gorm:"not null" json:"-"`
	User    User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
