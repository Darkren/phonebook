package repositories

import "github.com/Darkren/phonebook/models"

// PhoneRepository is an interface which interacts with
// the persistent data storage of phones
type PhoneRepository interface {
	Get(ID int64) (*models.Phone, error)
	ListByUser(userID int64) ([]*models.Phone, error)
	Add(phone *models.Phone) (int64, error)
	Update(phone *models.Phone) error
	Delete(ID int64) error
}
