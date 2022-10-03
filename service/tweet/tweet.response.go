package _tweet

import (
	"go-sample/entity"
	_user "go-sample/service/user"
)

type TweetResponse struct {
	ID        int64              `json:"id"`
	TweetContent string             `json:"tweet_content"`
	User      _user.UserResponse `json:"user,omitempty"`
}

func NewTweetResponse(tweet entity.Tweet) TweetResponse {
	return TweetResponse{
		ID:        tweet.ID,
		TweetContent: tweet.Content,
		User:      _user.NewUserResponse(tweet.User),
	}
}

func NewTweetArrayResponse(tweets []entity.Tweet) []TweetResponse {
	tweetRes := []TweetResponse{}
	for _, v := range tweets {
		p := TweetResponse{
			ID:        v.ID,
			TweetContent: v.Content,
			User:      _user.NewUserResponse(v.User),
		}
		tweetRes = append(tweetRes, p)
	}
	return tweetRes
}
