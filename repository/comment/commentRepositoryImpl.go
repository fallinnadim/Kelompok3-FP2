package repository

import (
	"database/sql"
	"errors"
	"fp2/dto"
	"fp2/entity"
)

type CommentRepositoryImpl struct {
	Db *sql.DB
}

// Create implements PhotoRepository.
func (p *CommentRepositoryImpl) Create(cp dto.CreateCommentRequest) entity.Comment {
	var newComment = entity.Comment{}
	query := `
		INSERT INTO comments (user_id, photo_id, message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`
	p.Db.QueryRow(query, cp.User_Id, cp.Photo_Id, cp.Message, cp.Created_At, cp.Updated_At).Scan(&newComment.Id, &newComment.User_Id, &newComment.Photo_Id, &newComment.Message, &newComment.Created_At, &newComment.Updated_At)

	return newComment
}

// Delete implements PhotoRepository.
func (p *CommentRepositoryImpl) Delete(id int) {
	query := `
		DELETE FROM comments WHERE id = $1;
	`
	p.Db.Exec(query, id)
}

// FindAll implements PhotoRepository.
func (p *CommentRepositoryImpl) FindAll(userId int) []dto.AllCommentResponse {
	comments := []dto.AllCommentResponse{}
	query := `
		SELECT c.id, c.user_id, c.photo_id, c.message, c.created_at, c.updated_at, u.id, u.username, u.email, p.id, p.title, p.caption, p.photo_url, p.user_id
		FROM comments AS c
		JOIN users AS u ON c.user_id = u.id
		JOIN photos AS p ON c.photo_id = p.id
		WHERE c.user_id = $1;
	`
	rows, _ := p.Db.Query(query, userId)
	defer rows.Close()
	for rows.Next() {
		comment := dto.AllCommentResponse{}
		rows.Scan(&comment.Id, &comment.User_Id, &comment.Photo_Id, &comment.Message, &comment.Created_At, &comment.Updated_At, &comment.User.Id, &comment.User.Username, &comment.User.Email, &comment.Photo.Id, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.Photo_Url, &comment.Photo.User_Id)
		comments = append(comments, comment)
	}
	return comments
}

// FindById implements PhotoRepository.
func (p *CommentRepositoryImpl) FindById(id int) (entity.Comment, error) {
	var comment = entity.Comment{}
	query := `
		SELECT id, user_id, photo_id, message, created_at, updated_at FROM comments WHERE id = $1;
	`
	errQuery := p.Db.QueryRow(query, id).Scan(&comment.Id, &comment.User_Id, &comment.Photo_Id, &comment.Message, &comment.Created_At, &comment.Updated_At)
	if errQuery == sql.ErrNoRows {
		return comment, errors.New("Comment not found")
	}
	return comment, nil
}

// Update implements PhotoRepository.
func (p *CommentRepositoryImpl) Update(up dto.UpdateCommentRequest) entity.Comment {
	var updateComment = entity.Comment{}
	query := `
		UPDATE comments
		SET message = $1, updated_at = $2
		WHERE id = $3
		RETURNING *;
	`
	p.Db.QueryRow(query, up.Message, up.Updated_At, up.Id).Scan(&updateComment.Id, &updateComment.User_Id, &updateComment.Photo_Id, &updateComment.Message, &updateComment.Created_At, &updateComment.Updated_At)

	return updateComment
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{Db: db}
}
