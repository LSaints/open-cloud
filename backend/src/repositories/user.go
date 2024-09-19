package repositories

import (
	"backend/src/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository UserRepository) Create(user models.User) (uint64, error) {
	query := `
		INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id
	`

	var userID uint64
	err := repository.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (repository UserRepository) GetAll(param string) ([]models.User, error) {
	param = fmt.Sprintf("%%%s%%", param)

	result, err := repository.db.Query(
		`SELECT id, name, email, created_at FROM users WHERE name ILIKE $1`,
		param,
	)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var users []models.User

	for result.Next() {
		var user models.User

		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepository) GetByID(ID uint64) (models.User, error) {
	result, err := repository.db.Query(
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer result.Close()

	var user models.User
	if result.Next() {
		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository UserRepository) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, user.Password, ID); err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = $1",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}
