package database

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func DoesTodayHaveSnapshot(db *gorm.DB) bool {
	return DoesDayHaveSnapshot(db, datatypes.Date(time.Now()))
}

func DoesDayHaveSnapshot(db *gorm.DB, date datatypes.Date) bool {
	var snapshot models.RestaurantSnapshot

	db.Where("datetime = ?", date).Take(&snapshot)
	return snapshot.ID != 0
}

func UpsertScraped(db *gorm.DB, scraped models.RestaurantSnapshot) error {
	var result = db.FirstOrCreate(&scraped)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SelectTodaysSnapshot(db *gorm.DB) (models.RestaurantSnapshot, error) {
	var today = datatypes.Date(time.Now())
	var snapshot = models.RestaurantSnapshot{Date: today}
	var result = db.Preload("Restaurants.MenuItems").Preload(clause.Associations).Find(&snapshot)

	return snapshot, result.Error
}

func CreateUser(db *gorm.DB, username string, passwordHash string) error {
	var user = models.User{
		Username:     username,
		PasswordHash: passwordHash,
	}

	var result = db.Create(&user)
	return result.Error
}

func GetUser(db *gorm.DB, username string) (models.User, error) {
	var user = models.User{}

	var result = db.Where(&models.User{Username: username}).Take(&user)
	return user, result.Error
}
