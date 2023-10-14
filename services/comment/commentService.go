package services

import (
	request "fp2/data/request/comment"
	response "fp2/data/response/comment"
)

type CommentService interface {
	Post(p request.CreateCommentRequest) (response.CreatedCommentResponse, error)
	GetAll(userId int) []response.AllCommentResponse
	Update(p request.UpdateCommentRequest) (response.UpdatedCommentResponse, error)
	Delete(id int) error
}
