package repository

import (
	"database/sql"
	"fmt"

	"github.com/ashiruhabeeb/simpleTodoApp/entity"
)

// userRepo holds db object of type database/sql package
type userRepo struct {
	db *sql.DB
}

// NewUserRepo constructor creates a new instance of todoRepo object
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

// InsertUser creates a new user record in the users table
func (ur *userRepo) InsertUser(u entity.User)(int, error){
	err := ur.db.QueryRow(insertuser, u.Username, u.FullName, u.Email, u.Password, u.Address, u.Avatar, u.DOB).Scan(&u.UserID)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return u.UserID, nil
}

// GetUser fetch a single user record from the users table 
func (ur *userRepo) GetUser(user_id int)(*entity.User, error){
	user := entity.User{}

	row := ur.db.QueryRow(getUser, user_id)

	err := row.Scan(&user.UserID, &user.Username, &user.FullName, &user.Email, &user.Address, &user.Avatar, &user.DOB, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == row.Err() {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

// GetUsers returns records in the users tables in batches of 3 ordered by their id
func (ur *userRepo) GetUsers()([]entity.User, error){
	rows, err := ur.db.Query(getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		var user entity.User
		rows.Scan(&user.UserID, &user.Username, &user.FullName, &user.Email, &user.Address, &user.Avatar, &user.DOB, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateTodo updates a single todo record in the todo table
func (ur *userRepo) UpdateUser(user_id int, username, fullname, address, dob string) error {
	_, err := ur.db.Exec(updateUser, username, fullname, address, dob)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a single todo record from the todo table
func (ur *userRepo) DeleteUser(user_id int) error {
	_, err := ur.db.Exec(deleteUser, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no record associated with user_id parameter: %v", err)
		}
		return err
	}
	return nil
}
