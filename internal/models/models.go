package models

import "gorm.io/datatypes"

type RestaurantSnapshot struct {
	ID   uint           `gorm:"primaryKey"`
	Date datatypes.Date `gorm:"uniqueIndex,sort:desc"`

	Restaurants []Restaurant `gorm:"foreignKey:RestaurantSnapshotID"`
}

type Restaurant struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"index"`
	Votes uint   `gorm:"default:0"`

	RestaurantSnapshotID uint
	MenuItems            []MenuItem `gorm:"foreignKey:RestaurantID"`
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
