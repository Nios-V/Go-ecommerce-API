package models

type User struct {
	Base
	Firstname string    `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname  string    `gorm:"type:varchar(30);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Roles     []Role    `gorm:"many2many:user_roles;" json:"roles"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`
}
