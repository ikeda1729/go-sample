package _user

import "go-sample/entity"

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
	}
}
