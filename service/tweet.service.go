package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"go-sample/dto"
	"go-sample/entity"
	"go-sample/repo"

	"github.com/mashingan/smapping"

	_tweet "go-sample/service/tweet"
)

type TweetService interface {
	All(userID string) (*[]_tweet.TweetResponse, error)
	CreateTweet(tweetRequest dto.CreateTweetRequest, userID string) (*_tweet.TweetResponse, error)
	UpdateTweet(updateTweetRequest dto.UpdateTweetRequest, userID string) (*_tweet.TweetResponse, error)
	FindOneTweetByID(tweetID string) (*_tweet.TweetResponse, error)
	DeleteTweet(tweetID string, userID string) error
}

type tweetService struct {
	tweetRepo repo.TweetRepository
}

func NewTweetService(tweetRepo repo.TweetRepository) TweetService {
	return &tweetService{
		tweetRepo: tweetRepo,
	}
}

func (c *tweetService) All(userID string) (*[]_tweet.TweetResponse, error) {
	tweets, err := c.tweetRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := _tweet.NewTweetArrayResponse(tweets)
	return &prods, nil
}

func (c *tweetService) CreateTweet(tweetRequest dto.CreateTweetRequest, userID string) (*_tweet.TweetResponse, error) {
	tweet := entity.Tweet{}
	err := smapping.FillStruct(&tweet, smapping.MapFields(&tweetRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	tweet.UserID = id
	p, err := c.tweetRepo.InsertTweet(tweet)
	if err != nil {
		return nil, err
	}

	res := _tweet.NewTweetResponse(p)
	return &res, nil
}

func (c *tweetService) FindOneTweetByID(tweetID string) (*_tweet.TweetResponse, error) {
	tweet, err := c.tweetRepo.FindOneTweetByID(tweetID)

	if err != nil {
		return nil, err
	}

	res := _tweet.NewTweetResponse(tweet)
	return &res, nil
}

func (c *tweetService) UpdateTweet(updateTweetRequest dto.UpdateTweetRequest, userID string) (*_tweet.TweetResponse, error) {
	tweet, err := c.tweetRepo.FindOneTweetByID(fmt.Sprintf("%d", updateTweetRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if tweet.UserID != uid {
		return nil, errors.New("produk ini bukan milik anda")
	}

	tweet = entity.Tweet{}
	err = smapping.FillStruct(&tweet, smapping.MapFields(&updateTweetRequest))

	if err != nil {
		return nil, err
	}

	tweet.UserID = uid
	tweet, err = c.tweetRepo.UpdateTweet(tweet)

	if err != nil {
		return nil, err
	}

	res := _tweet.NewTweetResponse(tweet)
	return &res, nil
}

func (c *tweetService) DeleteTweet(tweetID string, userID string) error {
	tweet, err := c.tweetRepo.FindOneTweetByID(tweetID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", tweet.UserID) != userID {
		return errors.New("produk ini bukan milik anda")
	}

	c.tweetRepo.DeleteTweet(tweetID)
	return nil

}
