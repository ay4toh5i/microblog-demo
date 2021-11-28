package userprofile

import (
	"time"
	useraccount "microblog/microblog/domain/entity/user_account"
)

type UserProfile struct {
	ID          int64
	UserId      useraccount.UserID `db:"user_id"`
	DisplayName string             `db:"display_name"`
	Description string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewUserProfile(userId int64, displayName string, description string) (*UserProfile, error) {
	return &UserProfile{
		UserId:      useraccount.UserID(userId),
		DisplayName: displayName,
		Description: description,
	}, nil
}
