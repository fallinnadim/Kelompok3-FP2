package services

import "fp2/dto"

type CommentService interface {
	Post(p dto.CreateCommentRequest) (dto.CreatedCommentResponse, error)
	GetAll(userId int) []dto.AllCommentResponse
	Update(p dto.UpdateCommentRequest) (dto.UpdatedCommentResponse, error)
	Delete(id int) error
}
