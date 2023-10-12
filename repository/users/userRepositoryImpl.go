package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"fp2/data/request/users"
	"fp2/models"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(id int) (user models.User, err error) {
	query := `
			SELECT * FROM users WHERE id = $1;
		`
	errQuery := u.Db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_At, &user.Updated_At)
	if errQuery == sql.ErrNoRows {
		return user, errors.New("user id not found")
	}
	return user, nil
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(id int) {
	query := `
		DELETE FROM users WHERE id = $1;
	`
	u.Db.Exec(query, id)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user request.UpdateUserRequest) models.User {
	fmt.Println(user)
	var updatedResult = models.User{}
	query := `
		UPDATE users
		SET username = $1, email = $2, updated_at = $3
		WHERE id = $4
		RETURNING *;
	`
	u.Db.QueryRow(query, user.Username, user.Email, user.Updated_At, user.Id).Scan(&updatedResult.Id, &updatedResult.Username, &updatedResult.Email, &updatedResult.Password, &updatedResult.Age, &updatedResult.Created_At, &updatedResult.Updated_At)

	return updatedResult
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
