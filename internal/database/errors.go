package database

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

const UniqueViolationErrorCode = "23505"
const DatabaseAlreadyExistsErrorCode = "42P04"

func IsUniqueViolation(err error) bool {
	var pgError *pgconn.PgError
	return errors.As(err, &pgError) && pgError.Code == UniqueViolationErrorCode
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

type DayAlreadyPopulatedError struct {
	day datatypes.Date
}

func (e *DayAlreadyPopulatedError) Error() string {
	return "The day " + time.Time(e.day).Format("2006-01-02") + "is already populated"
}
