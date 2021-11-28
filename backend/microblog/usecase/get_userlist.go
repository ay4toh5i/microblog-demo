package usecase

import "microblog/microblog/domain/queryservice/dto"

type GetUserListInput struct {
	CurrentUserID int64
}

type GetUserList func(input *GetUserListInput, presenter func(users []*dto.User)) error
