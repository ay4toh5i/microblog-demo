package dto

import (
	"time"
)

type Tweet struct {
	ID   string `json:"id"`
	User *struct {
		ID          int64  `json:"id"`
		Username    string `json:"username"`
		DisplayName string `db:"display_name" json:"display_name"`
		Description string `json:"description"`
	} `db:"user" json:"user"`
	Message   string    `json:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
