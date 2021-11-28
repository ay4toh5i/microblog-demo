package infra

import (
	"microblog/microblog/domain/entity/tweet"

	"github.com/jmoiron/sqlx"
)

type TweetRepo struct {
	db *sqlx.DB
}

func NewTweetRepo(db *sqlx.DB) *TweetRepo {
	return &TweetRepo{db}
}

func (r *TweetRepo) FindAll() ([]*tweet.Tweet, error) {
	var tweets []*tweet.Tweet

	err := r.db.Select(&tweets, "select * from tweets")
	if err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *TweetRepo) Create(t *tweet.Tweet) (*tweet.Tweet, error) {
	_, err := r.db.NamedExec("insert into tweets (id, user_id, message) values (:id, :user_id, :message)", t)
	if err != nil {
		return nil, err
	}

	var createdTweet tweet.Tweet
	err = r.db.Get(&createdTweet, "select * from tweets where id = ?", t.ID)
	if err != nil {
		return nil, err
	}

	return &createdTweet, nil
}
