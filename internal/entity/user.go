package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         string      `json:"id" gorm:"primaryKey"`
	FirstName  string      `json:"first_name" gorm:"not null"`
	LastName   string      `json:"last_name" gorm:"not null"`
	Age        int         `json:"age" gorm:"not null; check:age > 18"`
	Email      string      `json:"email" gorm:"not null; unique_index"`
	Activities []*Activity `json:"activities,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime:milli" swaggertype:"primitive,string"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoCreateTime:milli" swaggertype:"primitive,string"`
	DeletedAt  *time.Time  `json:"deleted_at,omitempty" gorm:"index" swaggertype:"primitive,string"`
}

func CreateUser(firstName string, lastName string, age int, email string) *User {
	return &User{
		ID:        strings.ToLower(uuid.New().String()),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Email:     email,
	}
}
