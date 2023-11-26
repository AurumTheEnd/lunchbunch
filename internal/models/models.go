package models

import "gorm.io/datatypes"

type RestaurantSnapshot struct {
	ID   uint           `gorm:"primaryKey"`
	Date datatypes.Date `gorm:"uniqueIndex,sort:desc"`

	CreatorID uint
	Creator   User

	Restaurants []Restaurant `gorm:"foreignKey:RestaurantSnapshotID"`
}

type Restaurant struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"index"`
	VotedOn bool   `gorm:"default:false"`

	RestaurantSnapshotID uint
	MenuItems            []MenuItem `gorm:"foreignKey:RestaurantID"`
	Votes                []Vote     `gorm:"foreignKey:RestaurantID"`
}

type MenuItem struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"index"`
	Price int

	RestaurantID uint
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex"`
	PasswordHash string
}

type Vote struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint
	UserID       uint
}
