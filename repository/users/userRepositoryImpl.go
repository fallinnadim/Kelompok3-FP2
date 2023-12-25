package repository

import (
	"database/sql"
	"errors"
	"fp2/dto"
	"fp2/entity"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(id int) (user entity.User, err error) {
	query := `
			SELECT id, username, email, password, age, created_at, updated_at FROM users WHERE id = $1;
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
func (u *UserRepositoryImpl) Update(user dto.UpdateUserRequest) entity.User {
	var updatedResult = entity.User{}
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
