package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Sport string

type Activity struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	Sport        Sport         `json:"sport" gorm:"not null"`
	Distance     float64       `json:"distance" gorm:"not null"`
	TrainingDate time.Time     `json:"training_date" swaggertype:"primitive,string" gorm:"not null"`
	Duration     time.Duration `json:"duration" swaggertype:"primitive,integer" gorm:"not null"`
	UserID       string        `json:"user_id"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime:milli" swaggertype:"primitive,string"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"autoCreateTime:milli" swaggertype:"primitive,string"`
	DeletedAt    *time.Time    `json:"deleted_at,omitempty" gorm:"index" swaggertype:"primitive,string"`
}

func CreateActivity(userId string, sport string, distance float64, trainigDate time.Time, duration int) *Activity {
	return &Activity{
		ID:           strings.ToLower(string(uuid.New().String())),
		Sport:        Sport(sport),
		Distance:     distance,
		TrainingDate: trainigDate,
		Duration:     time.Duration(duration),
		UserID:       userId,
	}
}
