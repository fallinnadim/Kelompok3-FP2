package repository

import (
	"database/sql"
	"errors"
	"fp2/data/request"
	"fp2/models"
	"log"
	"time"
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
	_, errQuery := u.Db.Exec(query, id)
	if errQuery != nil {
		log.Fatalln(errQuery.Error())
	}
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user models.User) {
	var updateUser = request.UpdateUserRequest{
		Id:         user.Id,
		Username:   user.Username,
		Email:      user.Email,
		Updated_At: time.Now().Format("2006-01-02"),
	}
	query := `
		UPDATE users
		SET username = $1, email = $2, updated_at = $3
		WHERE id = $4;
	`
	_, errQuery := u.Db.Exec(query, updateUser.Username, updateUser.Email, updateUser.Updated_At, updateUser.Id)
	if errQuery != nil {
		log.Fatalln(errQuery.Error())
	}
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}
