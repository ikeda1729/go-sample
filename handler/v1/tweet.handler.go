package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"go-sample/common/obj"
	"go-sample/common/response"
	"go-sample/dto"
	"go-sample/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TweetHandler interface {
	All(ctx *gin.Context)
	CreateTweet(ctx *gin.Context)
	UpdateTweet(ctx *gin.Context)
	DeleteTweet(ctx *gin.Context)
	FindOneTweetByID(ctx *gin.Context)
	FindTweetsByUserID(ctx *gin.Context)
}

type tweetHandler struct {
	tweetService service.TweetService
	jwtService   service.JWTService
}

func NewTweetHandler(tweetService service.TweetService, jwtService service.JWTService) TweetHandler {
	return &tweetHandler{
		tweetService: tweetService,
		jwtService:   jwtService,
	}
}

func (c *tweetHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	tweets, err := c.tweetService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", tweets)
	ctx.JSON(http.StatusOK, response)
}

func (c *tweetHandler) FindTweetsByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	tweets, err := c.tweetService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", tweets)
	ctx.JSON(http.StatusOK, response)
}

func (c *tweetHandler) CreateTweet(ctx *gin.Context) {
	var createTweetReq dto.CreateTweetRequest
	err := ctx.ShouldBind(&createTweetReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.tweetService.CreateTweet(createTweetReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *tweetHandler) FindOneTweetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.tweetService.FindOneTweetByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *tweetHandler) DeleteTweet(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.tweetService.DeleteTweet(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *tweetHandler) UpdateTweet(ctx *gin.Context) {
	updateTweetRequest := dto.UpdateTweetRequest{}
	err := ctx.ShouldBind(&updateTweetRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateTweetRequest.ID = id
	tweet, err := c.tweetService.UpdateTweet(updateTweetRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", tweet)
	ctx.JSON(http.StatusOK, response)

}
