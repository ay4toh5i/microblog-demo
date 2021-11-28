package tweet

type TweetRepository interface {
	Create(*Tweet) (*Tweet, error)
}
