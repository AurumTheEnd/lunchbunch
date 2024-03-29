package database

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/server/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func dayBoundaries(t time.Time) (dayStart time.Time, dayEnd time.Time) {
	dayStart = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	dayEnd = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return
}

func CreateScraped(db *gorm.DB, scraped *models.RestaurantSnapshot) error {
	var result = db.Create(scraped).
		Clauses(clause.Returning{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SelectSnapshotWith(db *gorm.DB, id uint) (models.RestaurantSnapshot, error) {
	var snapshot models.RestaurantSnapshot

	var result = db.
		Preload("Restaurants.MenuItems").
		Preload("Restaurants.Votes").
		Preload(clause.Associations).
		Where("id = ?", id).
		Take(&snapshot)

	return snapshot, result.Error
}

func UpdateVotedOn(db *gorm.DB, formResult *data.NewPollFormDataToServer) error {
	var tx = db.Begin()

	tx.Model(&models.Restaurant{}).
		Where("restaurant_snapshot_id = ?", formResult.SnapshotID).
		Where("id IN ?", formResult.Checked).
		Update("voted_on", "TRUE")

	tx.Model(&models.RestaurantSnapshot{}).
		Where("id = ?", formResult.SnapshotID).
		Update("has_poll_started", true)

	var result = tx.Commit()

	return result.Error
}

func SelectSnapshotById(db *gorm.DB, id uint) (models.RestaurantSnapshot, error) {
	var snapshots models.RestaurantSnapshot

	var result = db.
		Preload("Restaurants", func(db *gorm.DB) *gorm.DB {
			return db.Where("voted_on = ?", true).
				Order("restaurants.name ASC")
		}).
		Preload("Restaurants.MenuItems").
		Preload("Restaurants.Votes").
		Preload("Restaurants.Votes.User", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password_hash")
		}).
		Preload(clause.Associations).
		Where("id = ?", id).
		Take(&snapshots)

	return snapshots, result.Error
}

func UpdateVote(db *gorm.DB, restaurantId uint, userId uint, shouldCast bool) error {
	var resultError error
	if shouldCast {
		resultError = db.Create(&models.Vote{
			RestaurantID: restaurantId,
			UserID:       userId,
		}).Error
	} else {
		resultError = db.Delete(&models.Vote{}, "restaurant_id = ? AND user_id = ?", restaurantId, userId).Error
	}

	return resultError
}

func SelectTodaysSnapshots(db *gorm.DB) ([]models.RestaurantSnapshot, error) {
	var timestamp = time.Now()
	var dayStart, dayEnd = dayBoundaries(timestamp)
	var snapshots []models.RestaurantSnapshot

	var result = db.
		Preload("Restaurants", func(db *gorm.DB) *gorm.DB {
			return db.Where("voted_on = ?", true).
				Order("restaurants.name ASC")
		}).
		Preload("Restaurants.MenuItems").
		Preload("Restaurants.Votes").
		Preload(clause.Associations).
		Where("datetime >= ?", dayStart).
		Where("datetime <= ?", dayEnd).
		Where("has_poll_started = TRUE").
		Order("datetime DESC").
		Find(&snapshots)

	return snapshots, result.Error
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
