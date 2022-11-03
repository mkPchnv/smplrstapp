package dto

type ActivityCreateDto struct {
	UserId       string  `json:"user_id" mapstructure:"user_id"`
	Sport        string  `json:"sport"`
	Duration     int     `json:"duration"`
	TrainingDate string  `json:"training_date" mapstructure:"training_date"`
	Distance     float64 `json:"distance"`
}
