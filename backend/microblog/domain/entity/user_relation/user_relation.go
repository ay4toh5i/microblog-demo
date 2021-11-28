package userrelation

import useraccount "microblog/microblog/domain/entity/user_account"

type UserRelation struct {
	FollowerUserID useraccount.UserID `db:"follower_user_id"`
	FolloweeUserId useraccount.UserID `db:"followee_user_id"`
}
