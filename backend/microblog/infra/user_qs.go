package infra

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice/dto"

	"github.com/jmoiron/sqlx"
)

type UserQS struct {
	db *sqlx.DB
}

func NewUserQS(db *sqlx.DB) *UserQS {
	return &UserQS{db}
}

func (q *UserQS) FindAll(followerUserID useraccount.UserID) ([]*dto.User, error) {
	var users []*dto.User

	err := q.db.Select(&users,
		`select 
      ua.id,
      ua.username,
      up.display_name,
      up.description,
      case when following.id is null then 0 else 1 end as "following",
      case when followed.id is null then 0 else 1 end as "followed_by"
     from user_accounts ua 
       inner join user_profiles up 
         on ua.id = up.user_id
       left join user_relations following
         on ua.id = following.followee_user_id and following.follower_user_id = ?
       left join user_relations followed
         on  ua.id = followed.follower_user_id and followed.followee_user_id = ?`,
		followerUserID,
		followerUserID,
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}
