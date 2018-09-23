// Package mem contains a UserRepository interface implementation
// which utilizes slice as in-memory data storage
package mem

import (
	"sync"

	"github.com/Darkren/phonebook/models"
	"github.com/Darkren/phonebook/repositories"
)

// UserRepository is a UserRepository inetrface implementation
type UserRepository struct {
	users  []*models.User
	serial int64
	sync.RWMutex
}

// New constructs new UserRepository
func New() repositories.UserRepository {
	return &UserRepository{}
}

// Get gets user by its ID
func (r *UserRepository) Get(ID int64) (*models.User, error) {
	r.RLock()
	defer r.RUnlock()

	for _, user := range r.users {
		if user.ID == ID {
			return user, nil
		}
	}

	return nil, nil
}

// List gets all added users
func (r *UserRepository) List() ([]*models.User, error) {
	r.RLock()
	defer r.RUnlock()

	users := []*models.User{}

	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Add adds user to the repository
func (r *UserRepository) Add(user *models.User) (int64, error) {
	r.Lock()
	defer r.Unlock()

	r.serial++
	user.ID = r.serial

	r.users = append(r.users, user)

	return user.ID, nil
}

// Update updates already stored user by its ID
func (r *UserRepository) Update(user *models.User) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.users {
		if r.users[i].ID == user.ID {
			r.users[i] = user

			break
		}
	}

	return nil
}

// Delete removes user by ID from repository
func (r *UserRepository) Delete(ID int64) error {
	r.Lock()
	defer r.Unlock()

	for i := range r.users {
		if r.users[i].ID == ID {
			r.users = append(r.users[:i], r.users[i+1:]...)

			break
		}
	}

	return nil
}
