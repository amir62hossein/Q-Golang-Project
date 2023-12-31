package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"` // UserID in other Tables
	Name        string      `json:"name"`
	DateOfBirth time.Time   `json:"date_of_birth"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	Passengers  []uuid.UUID `gorm:"type:uuid[];Column:passengers" json:"passengers"`
	CreatedAt   time.Time   `json:"created_at"`
	ModifiedAt  time.Time   `json:"modified_at"`
	DeletedAt   time.Time   `json:"deleted_at"`
}
