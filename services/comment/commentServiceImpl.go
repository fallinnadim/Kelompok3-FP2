package services

import (
	"errors"
	request "fp2/data/request/comment"
	response "fp2/data/response/comment"
	"fp2/helper"
	repository "fp2/repository/comment"
	pRepository "fp2/repository/photo"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	PhotoRepository   pRepository.PhotoRepository
	Validate          *validator.Validate
}

// Delete implements PhotoService.
func (c *CommentServiceImpl) Delete(id int) error {
	// Panggil service
	c.CommentRepository.Delete(id)
	return nil
}

// GetAll implements PhotoService.
func (c *CommentServiceImpl) GetAll(userId int) []response.AllCommentResponse {
	result := c.CommentRepository.FindAll(userId)
	return result
}

// Post implements PhotoService.
func (c *CommentServiceImpl) Post(cp request.CreateCommentRequest) (response.CreatedCommentResponse, error) {
	// Fungsi cek apakah foto beneran ada atau tidak
	_, errCheck := c.PhotoRepository.FindById(cp.Photo_Id)
	if errCheck != nil {
		return response.CreatedCommentResponse{}, &helper.RequestError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("Photo Not Found"),
		}
	}
	// Validasi Struct
	errValidation := c.Validate.Struct(cp)
	if errValidation != nil {
		return response.CreatedCommentResponse{}, errValidation
	}
	cp.Created_At = time.Now().Format("2006-01-02")
	cp.Updated_At = time.Now().Format("2006-01-02")
	// Panggil Repository
	result := c.CommentRepository.Create(cp)
	// Return
	resp := response.CreatedCommentResponse{
		Id:         result.Id,
		Message:    result.Message,
		Photo_Id:   result.Photo_Id,
		User_Id:    result.User_Id,
		Created_At: result.Created_At,
	}
	return resp, nil
}

// Update implements PhotoService.
func (c *CommentServiceImpl) Update(cp request.UpdateCommentRequest) (response.UpdatedCommentResponse, error) {
	// Validasi Struct
	errValidation := c.Validate.Struct(cp)
	if errValidation != nil {
		return response.UpdatedCommentResponse{}, errValidation
	}
	cp.Updated_At = time.Now().Format("2006-01-02")
	// Panggil service
	result := c.CommentRepository.Update(cp)
	updateComment := response.UpdatedCommentResponse{
		Id:         result.Id,
		Message:    result.Message,
		Photo_Id:   result.Photo_Id,
		User_Id:    result.User_Id,
		Updated_At: result.Updated_At,
	}
	return updateComment, nil
}

func NewCommentServiceImpl(c repository.CommentRepository, p pRepository.PhotoRepository, v *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: c,
		PhotoRepository:   p,
		Validate:          v,
	}
}
