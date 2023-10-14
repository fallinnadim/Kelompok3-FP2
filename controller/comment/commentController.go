package controller

import (
	request "fp2/data/request/comment"
	response "fp2/data/response/users"
	"fp2/helper"
	services "fp2/services/comment"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentService services.CommentService
}

func NewCommentController(s services.CommentService) *CommentController {
	return &CommentController{CommentService: s}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid token",
		})
		return
	}
	// panggil service
	createComment := request.CreateCommentRequest{}
	err := ctx.ShouldBindJSON(&createComment)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	createComment.User_Id = userId.(int)
	result, errCreate := c.CommentService.Post(createComment)
	// return response
	if errCreate != nil {
		statusCode, errMessage := helper.ParseError(errCreate)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *CommentController) GetAllComment(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	webResponse := c.CommentService.GetAll(userId.(int))
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	// panggil service
	updateCommentRequest := request.UpdateCommentRequest{}
	err := ctx.ShouldBindJSON(&updateCommentRequest)
	if err != nil {
		statusCode, errMessage := helper.ParseError(err)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	updateCommentRequest.Id = commentId
	result, errUpdate := c.CommentService.Update(updateCommentRequest)
	// return response
	if errUpdate != nil {
		statusCode, errMessage := helper.ParseError(errUpdate)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	// panggil service
	errDelete := c.CommentService.Delete(commentId)
	// return response
	if errDelete != nil {
		statusCode, errMessage := helper.ParseError(errDelete)
		webResponse := response.FailedResponse{
			Status:  false,
			Message: errMessage,
		}
		ctx.JSON(statusCode, webResponse)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comments has been successfully deleted",
	})
}
