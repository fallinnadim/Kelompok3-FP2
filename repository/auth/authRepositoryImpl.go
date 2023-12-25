package repository

import (
	"database/sql"
	"errors"
	"fp2/entity"
)

type AuthRepositoryImpl struct {
	Db *sql.DB
}

// Register implements UserRepository.
func (a *AuthRepositoryImpl) Create(user entity.User) entity.User {
	var newUser = entity.User{}
	query := `
		INSERT INTO users (username, email, password, age, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`
	a.Db.QueryRow(query, user.Username, user.Email, user.Password, user.Age, user.Created_At, user.Updated_At).Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password, &newUser.Age, &newUser.Created_At, &newUser.Updated_At)

	return newUser

}

// Find Email.
func (a *AuthRepositoryImpl) FindEmail(email string) (user entity.User, err error) {
	query := `
			SELECT id, username, email, password, age, created_at, updated_at FROM users WHERE email = $1;
		`
	errQuery := a.Db.QueryRow(query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_At, &user.Updated_At)
	if errQuery == sql.ErrNoRows {
		return user, errors.New("email not found")
	}
	return user, nil
}

// Find Username.
func (a *AuthRepositoryImpl) FindUsername(username string) (user entity.User, err error) {
	query := `
			SELECT id, username, email, password, age, created_at, updated_at FROM users WHERE username = $1;
		`
	errQuery := a.Db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_At, &user.Updated_At)
	if errQuery == sql.ErrNoRows {
		return user, errors.New("username not found")
	}
	return user, nil
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}
