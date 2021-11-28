package infra

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice/dto"

	"github.com/jmoiron/sqlx"
)

type TweetQS struct {
	db *sqlx.DB
}

func NewTweetQS(db *sqlx.DB) *TweetQS {
	return &TweetQS{db}
}

func (q *TweetQS) FindAll(followerUserID useraccount.UserID) ([]*dto.Tweet, error) {
	var tweets []*dto.Tweet

	err := q.db.Select(
		&tweets,
		`select 
       t.id as "id",
       t.message as "message",
       t.created_at as "created_at",
       t.updated_at as "updated_at",
       ua.id as "user.id",
       ua.username as "user.username",
       up.display_name as "user.display_name",
       up.description as "user.description"
     from tweets t 
       inner join user_accounts ua 
         on t.user_id = ua.id
       inner join user_profiles up
         on t.user_id = up.user_id
       inner join user_relations ur
         on t.user_id = ur.followee_user_id
       where ur.follower_user_id = ?
      order by id desc `,
		followerUserID,
	)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
