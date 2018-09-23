// Package mem contains the mem repository factory
package mem

import (
	"github.com/Darkren/phonebook/repositories"
	"github.com/Darkren/phonebook/repositories/fabric"
	phoneRepository "github.com/Darkren/phonebook/repositories/phone/mem"
	userRepository "github.com/Darkren/phonebook/repositories/user/mem"
)

// RepositoryFabric is a fabric of mem repositories
type RepositoryFabric struct {
}

// New constructs the fabric instance
func New() fabric.RepositoryFabric {
	return &RepositoryFabric{}
}

// CreateUserRepository constructs mem user repository
func (f *RepositoryFabric) CreateUserRepository() repositories.UserRepository {
	return userRepository.New()
}

// CreatePhoneRepository constructs mem phone repository
func (f *RepositoryFabric) CreatePhoneRepository() repositories.PhoneRepository {
	return phoneRepository.New()
}
