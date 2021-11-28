package queryservice

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice/dto"
)

type UserQueryService interface {
	FindAll(followerUserID useraccount.UserID) ([]*dto.User, error)
}
