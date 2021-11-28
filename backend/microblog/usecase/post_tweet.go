package usecase

type PostTweetInputData struct {
	CurrentUserID int64
	Message       string
}

type PostTweet func(input *PostTweetInputData, presenter func()) error
