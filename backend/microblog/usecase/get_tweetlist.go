package usecase

import "microblog/microblog/domain/queryservice/dto"

type GetTweetListInput struct {
	CurrentUserID int64
}

type GetTweetList func(input *GetTweetListInput, presenter func(tweets []*dto.Tweet)) error
