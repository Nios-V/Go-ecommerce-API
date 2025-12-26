package models

type Role struct {
	Name        string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Description *string `gorm:"type:varchar(255)" json:"description,omitempty"`

	Base
}
