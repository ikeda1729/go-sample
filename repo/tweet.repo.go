package repo

import (
	"go-sample/entity"

	"gorm.io/gorm"
)

type TweetRepository interface {
	All(userID string) ([]entity.Tweet, error)
	InsertTweet(tweet entity.Tweet) (entity.Tweet, error)
	UpdateTweet(tweet entity.Tweet) (entity.Tweet, error)
	DeleteTweet(tweetID string) error
	FindOneTweetByID(ID string) (entity.Tweet, error)
	// FindAllTweet(userID string) ([]entity.Tweet, error)
}

type tweetRepo struct {
	connection *gorm.DB
}

func NewTweetRepo(connection *gorm.DB) TweetRepository {
	return &tweetRepo{
		connection: connection,
	}
}

func (c *tweetRepo) All(userID string) ([]entity.Tweet, error) {
	tweets := []entity.Tweet{}
	c.connection.Preload("User").Where("user_id = ?", userID).Order("created_at DESC").Find(&tweets)
	return tweets, nil
}

func (c *tweetRepo) InsertTweet(tweet entity.Tweet) (entity.Tweet, error) {
	c.connection.Save(&tweet)
	c.connection.Preload("User").Find(&tweet)
	return tweet, nil
}

func (c *tweetRepo) UpdateTweet(tweet entity.Tweet) (entity.Tweet, error) {
	c.connection.Save(&tweet)
	c.connection.Preload("User").Find(&tweet)
	return tweet, nil
}

func (c *tweetRepo) FindOneTweetByID(tweetID string) (entity.Tweet, error) {
	var tweet entity.Tweet
	res := c.connection.Preload("User").Where("id = ?", tweetID).Take(&tweet)
	if res.Error != nil {
		return tweet, res.Error
	}
	return tweet, nil
}

// func (c *tweetRepo) FindAllTweet(userID string) ([]entity.Tweet, error) {
// 	tweets := []entity.Tweet{}
// 	c.connection.Where("user_id = ?", userID).Find(&tweets)
// 	return tweets, nil
// }

func (c *tweetRepo) DeleteTweet(tweetID string) error {
	var tweet entity.Tweet
	res := c.connection.Preload("User").Where("id = ?", tweetID).Take(&tweet)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&tweet)
	return nil
}
