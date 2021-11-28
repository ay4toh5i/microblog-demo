package interactor

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice"
	"microblog/microblog/domain/queryservice/dto"
	"microblog/microblog/usecase"
)

func GetTweetList(tqs queryservice.TweetQueryService) usecase.GetTweetList {
	return func(input *usecase.GetTweetListInput, presenter func(tweets []*dto.Tweet)) error {
		userID := useraccount.UserID(input.CurrentUserID)

		tweets, err := tqs.FindAll(userID)
		if err != nil {
			return err
		}

		presenter(tweets)

		return nil
	}
}
