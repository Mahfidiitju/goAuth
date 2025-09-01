package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"time"
)

type UserRepository interface {
	GetByID() (*models.User, error)
	Create(username string, email string, password string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteByID(id int64) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {

	return []*models.User{}, nil
}

func (u *UserRepositoryImpl) DeleteByID(id int64) error {
	return nil
}

func (u *UserRepositoryImpl) Create(username string, email string, password string) (*models.User, error) {
	now := time.Now().Format(time.RFC3339)
	query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	result, err := u.db.Exec(query, username, email, password, now, now)
	if err != nil {
		return nil, err
	}

	id, rowErr := result.LastInsertId()
	if rowErr != nil {
		return nil, rowErr
	}

	user := &models.User{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Getching user in UserRepository")

	// Step 1: Prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, 1)

	// Step 3: Process the result
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	// Step 4: Print the user details
	fmt.Println("User fetched successfully:", user)

	return user, nil
}
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	row := u.db.QueryRow(query, email)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}
	return user, nil
}
