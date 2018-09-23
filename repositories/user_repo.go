package repositories

import "github.com/Darkren/phonebook/models"

// UserRepository is an interface which interacts with
// the persistent data storage of users
type UserRepository interface {
	Get(ID int64) (*models.User, error)
	List() ([]*models.User, error)
	Add(user *models.User) (int64, error)
	Update(user *models.User) error
	Delete(ID int64) error
}
