package _tweet

import (
	"go-sample/entity"
	_user "go-sample/service/user"
)

type TweetResponse struct {
	ID           int64              `json:"id"`
	TweetContent string             `json:"tweet_content"`
	CreatedAt    string             `json:"created_at"`
	User         _user.UserResponse `json:"user,omitempty"`
}

func NewTweetResponse(tweet entity.Tweet) TweetResponse {
	return TweetResponse{
		ID:           tweet.ID,
		TweetContent: tweet.Content,
		CreatedAt:    tweet.CreatedAt.String(),
		User:         _user.NewUserResponse(tweet.User),
	}
}

func NewTweetArrayResponse(tweets []entity.Tweet) []TweetResponse {
	tweetRes := []TweetResponse{}
	for _, v := range tweets {
		p := TweetResponse{
			ID:           v.ID,
			TweetContent: v.Content,
			CreatedAt:    v.CreatedAt.String(),
			User:         _user.NewUserResponse(v.User),
		}
		tweetRes = append(tweetRes, p)
	}
	return tweetRes
}
