package contact

import (
	"time"
)

type AddContact struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateContact struct {
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}
