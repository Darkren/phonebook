// Package mem contains a PhoneRepository interface implementation
// which utilizes slice as in-memory data storage
package mem

import (
	"sync"

	"github.com/Darkren/phonebook/models"
	"github.com/Darkren/phonebook/repositories"
)

// PhoneRepository is a PhoneRepository inetrface implementation
type PhoneRepository struct {
	phones []*models.Phone
	serial int64
	sync.RWMutex
}

// New constructs new PhoneRepository
func New() repositories.PhoneRepository {
	return &PhoneRepository{}
}

// Get gets phone by its ID
func (r *PhoneRepository) Get(ID int64) (*models.Phone, error) {
	r.RLock()
	defer r.RUnlock()

	for _, phone := range r.phones {
		if phone.ID == ID {
			return phone, nil
		}
	}

	return nil, nil
}

// ListByUser gets phones list by their user ID
func (r *PhoneRepository) ListByUser(userID int64) ([]*models.Phone, error) {
	r.RLock()
	defer r.RUnlock()

	phones := []*models.Phone{}

	for _, phone := range r.phones {
		if phone.UserID == userID {
			phones = append(phones, phone)
		}
	}

	return phones, nil
}

// Add adds phone to the repository
func (r *PhoneRepository) Add(phone *models.Phone) (int64, error) {
	r.Lock()
	defer r.Unlock()

	r.serial++
	phone.ID = r.serial

	r.phones = append(r.phones, phone)

	return phone.ID, nil
}

// Update updates already stored phone by its ID
func (r *PhoneRepository) Update(phone *models.Phone) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.phones {
		if r.phones[i].ID == phone.ID {
			r.phones[i] = phone

			break
		}
	}

	return nil
}

// Delete removes phone by ID from repository
func (r *PhoneRepository) Delete(ID int64) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.phones {
		if r.phones[i].ID == ID {
			r.phones = append(r.phones[:i], r.phones[i+1:]...)

			break
		}
	}

	return nil
}
