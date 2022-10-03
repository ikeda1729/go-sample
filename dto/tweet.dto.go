package dto

type CreateTweetRequest struct {
	Content string `json:"content" form:"name" binding:"required,min=1"`
}

type UpdateTweetRequest struct {
	ID   int64  `json:"id" form:"id"`
	Content string `json:"content" form:"name" binding:"required,min=1"`
}
