package useraccount

import (
	"time"
)

type UserAccount struct {
	ID        UserID
	Username  Username
	Password  *Password
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserAccount(username, password string) (*UserAccount, error) {
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &UserAccount{
		Username: Username(username),
		Password: p,
	}, nil
}

func (u *UserAccount) VerifyPassword(password string) bool {
	return u.Password.CheckPasswordHash(password)
}
