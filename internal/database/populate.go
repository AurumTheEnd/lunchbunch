package database

import (
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"gorm.io/gorm"
)

func Populate(db *gorm.DB, scraped models.RestaurantSnapshot) (err error) {
	if DoesDayHaveSnapshot(db, scraped.Date) {
		err = &DayAlreadyPopulatedError{day: scraped.Date}
		return
	}

	var result = db.FirstOrCreate(&scraped)
	if result.Error != nil {
		err = result.Error
	}

	return
}
