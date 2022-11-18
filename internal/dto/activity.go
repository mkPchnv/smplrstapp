package dto

import "smplrstapp/internal/entity"

type ActivityCreateDto struct {
	UserId       string       `json:"user_id" mapstructure:"user_id"`
	Sport        entity.Sport `json:"sport" enums:"run,swim,bike" example:"bike"`
	Duration     int          `json:"duration"`
	TrainingDate string       `json:"training_date" mapstructure:"training_date"`
	Distance     float64      `json:"distance"`
}
