package models

type User struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"size:50;not null"`
	Age    int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID                uint   `gorm:"primaryKey"`
	UserID            uint   `gorm:"not null;unique"`
	Bio               string `gorm:"size:255"`
	ProfilePictureURL string `gorm:"size:255"`
}
