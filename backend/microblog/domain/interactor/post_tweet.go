package interactor

import (
	"microblog/microblog/domain/entity/tweet"
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/usecase"
)

func PostTweet(tr tweet.TweetRepository) usecase.PostTweet {
	return func(input *usecase.PostTweetInputData, presenter func()) error {
		userID := useraccount.UserID(input.CurrentUserID)

		tweet, err := tweet.NewTweet(userID, input.Message)
		if err != nil {
			return err
		}

		_, err = tr.Create(tweet)
		if err != nil {
			return err
		}

		presenter()

		return nil
	}
}
