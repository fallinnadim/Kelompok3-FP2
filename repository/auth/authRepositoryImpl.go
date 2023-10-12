package repository

import (
	"database/sql"
	"errors"
	"fp2/models"
)

type AuthRepositoryImpl struct {
	Db *sql.DB
}

// Register implements UserRepository.
func (a *AuthRepositoryImpl) Create(user models.User) error {
	query := `
		INSERT INTO users (username, email, password, age, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, errQuery := a.Db.Exec(query, user.Username, user.Email, user.Password, user.Age, user.Created_At, user.Updated_At)
	if errQuery != nil {
		return errQuery
	}
	return nil
}

// Find Email.
func (a *AuthRepositoryImpl) FindEmail(username string) (user models.User, err error) {
	query := `
			SELECT * FROM users WHERE username = $1;
		`
	errQuery := a.Db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Age, &user.Created_At, &user.Updated_At)
	if errQuery == sql.ErrNoRows {
		return user, errors.New("user id not found")
	}
	return user, nil
}

func NewAuthRepositoryImpl(db *sql.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}
