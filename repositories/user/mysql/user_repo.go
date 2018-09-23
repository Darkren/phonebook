// Package mysql contains a PhoneRepository interface implementation
// which utilizes MySQL as persistent data storage
package mysql

import (
	"database/sql"
	"fmt"

	"github.com/Darkren/phonebook/models"
	"github.com/Darkren/phonebook/repositories"
)

const (
	tableName = "users"
)

// UserRepository is a UserRepository inetrface implementation
type UserRepository struct {
	db *sql.DB
}

// New constructs new UserRepository
func New(db *sql.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

// Get gets user by its ID
func (r *UserRepository) Get(ID int64) (*models.User, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?;", tableName)

	user := models.User{}

	err := r.db.QueryRow(sql, ID).Scan(&user.ID, &user.Name, &user.Surname, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// List gets all added users
func (r *UserRepository) List() ([]*models.User, error) {
	sql := fmt.Sprintf("SELECT * FROM %s;", tableName)

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// Add adds user to the repository
func (r *UserRepository) Add(user *models.User) (int64, error) {
	sql := fmt.Sprintf(`
		INSERT INTO %s 
			(name, surname, age) 
		VALUES 
			(?, ?, ?);`, tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(user.Name, user.Surname, user.Age)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Update updates already stored user by its ID
func (r *UserRepository) Update(user *models.User) error {
	sql := fmt.Sprintf(`
		UPDATE %s SET
			name = ?,
			surname = ?,
			age = ?
		WHERE 
			id = ?;`, tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Surname, user.Age, user.ID)

	return err
}

// Delete removes user by ID from repository
func (r *UserRepository) Delete(ID int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)

	return err
}
