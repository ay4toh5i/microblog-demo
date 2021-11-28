package queryservice

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice/dto"
)

type TweetQueryService interface {
	FindAll(followerUserID useraccount.UserID) ([]*dto.Tweet, error)
}
