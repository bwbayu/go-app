package repository

import (
	"database/sql"
	"go-app/internal/models"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (name, email, university, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.University, time.Now(), time.Now())
	return err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := "SELECT id, name, email, university, created_at, updated_at FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.University, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	query := "SELECT id, name, email, university, created_at, updated_at FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email,
		&user.University, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *UserRepository) Update(user *models.User) error {
	query := "UPDATE users SET name = ?, email = ?, university = ? WHERE id = ?"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.University, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
