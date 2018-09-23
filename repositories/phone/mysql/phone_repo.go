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
	tableName = "phones"
)

// PhoneRepository is a PhoneRepository inetrface implementation
type PhoneRepository struct {
	db *sql.DB
}

// New constructs new PhoneRepository
func New(db *sql.DB) repositories.PhoneRepository {
	return &PhoneRepository{db: db}
}

// Get gets phone by its ID
func (r *PhoneRepository) Get(ID int64) (*models.Phone, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?;", tableName)

	phone := models.Phone{}

	err := r.db.QueryRow(sql, ID).Scan(&phone.ID, &phone.UserID, &phone.Phone)
	if err != nil {
		return nil, err
	}

	return &phone, nil
}

// ListByUser gets phones list by their user ID
func (r *PhoneRepository) ListByUser(userID int64) ([]*models.Phone, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?;", tableName)

	rows, err := r.db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	phones := []*models.Phone{}

	for rows.Next() {
		phone := models.Phone{}
		err := rows.Scan(&phone.ID, &phone.Phone, &phone.UserID)
		if err != nil {
			return nil, err
		}

		phones = append(phones, &phone)
	}

	return phones, nil
}

// Add adds phone to the repository
func (r *PhoneRepository) Add(phone *models.Phone) (int64, error) {
	sql := fmt.Sprintf(`
		INSERT INTO %s
			(user_id, phone) 
		VALUES 
			(?, ?);`, tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(phone.UserID, phone.Phone)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Update updates already stored phone by its ID
func (r *PhoneRepository) Update(phone *models.Phone) error {
	sql := fmt.Sprintf(`
		UPDATE %s SET
			user_id = ?,
			phone = ?
		WHERE 
			id = ?;`, tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(phone.UserID, phone.Phone, phone.ID)

	return err
}

// Delete removes phone by ID from repository
func (r *PhoneRepository) Delete(ID int64) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)

	return err
}
