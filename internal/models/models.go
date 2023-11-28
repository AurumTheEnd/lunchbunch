package models

import (
	"time"
)

type RestaurantSnapshot struct {
	ID             uint      `gorm:"primaryKey"`
	Datetime       time.Time `gorm:"not null"`
	HasPollStarted bool      `gorm:"not null,default:false"`

	CreatorID uint `gorm:"not null"`
	Creator   User

	Restaurants []Restaurant `gorm:"foreignKey:RestaurantSnapshotID"`
}

type Restaurant struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null,index"`
	VotedOn bool   `gorm:"not null,default:false"`

	RestaurantSnapshotID uint
	MenuItems            []MenuItem `gorm:"foreignKey:RestaurantID"`
	Votes                []Vote     `gorm:"foreignKey:RestaurantID"`
}

type MenuItem struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null,index"`
	Price int

	RestaurantID uint
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null,uniqueIndex"`
	PasswordHash string `gorm:"not null"`
}

type Vote struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint `gorm:"not null"`

	UserID uint `gorm:"not null"`
	User   User
}
