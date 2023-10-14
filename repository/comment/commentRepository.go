package repository

import (
	request "fp2/data/request/comment"
	response "fp2/data/response/comment"
	"fp2/models"
)

type CommentRepository interface {
	FindAll(userId int) []response.AllCommentResponse
	Create(sm request.CreateCommentRequest) models.Comment
	Update(sm request.UpdateCommentRequest) models.Comment
	Delete(id int)
	FindById(id int) (models.Comment, error)
}
