package repository

import (
	"fp2/dto"
	"fp2/entity"
)

type CommentRepository interface {
	FindAll(userId int) []dto.AllCommentResponse
	Create(sm dto.CreateCommentRequest) entity.Comment
	Update(sm dto.UpdateCommentRequest) entity.Comment
	Delete(id int)
	FindById(id int) (entity.Comment, error)
}
