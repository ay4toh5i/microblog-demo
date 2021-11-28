package tweet

import (
	"time"
	useraccount "microblog/microblog/domain/entity/user_account"
)

type Tweet struct {
	ID        ULID
	UserId    useraccount.UserID `db:"user_id"`
	Message   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewTweet(userID useraccount.UserID, message string) (*Tweet, error) {
	return &Tweet{
		ID:      NewULID(),
		UserId:  userID,
		Message: message,
	}, nil
}
