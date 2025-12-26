package models

type Category struct {
	Name        string  `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Description *string `gorm:"type:varchar(255)" json:"description,omitempty"`

	Products []Product `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`

	Base
}
