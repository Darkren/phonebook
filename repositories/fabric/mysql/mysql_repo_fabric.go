// Package mysql contains the mysql repository factory
package mysql

import (
	"database/sql"

	"github.com/Darkren/phonebook/repositories"
	"github.com/Darkren/phonebook/repositories/fabric"
	phoneRepository "github.com/Darkren/phonebook/repositories/phone/mysql"
	userRepository "github.com/Darkren/phonebook/repositories/user/mysql"
)

// RepositoryFabric is a factory of mysql repositories
type RepositoryFabric struct {
	db *sql.DB
}

// New constructs mysql fabric instace
func New(db *sql.DB) fabric.RepositoryFabric {
	return &RepositoryFabric{db: db}
}

// CreateUserRepository constructs mysql user repository
func (f *RepositoryFabric) CreateUserRepository() repositories.UserRepository {
	return userRepository.New(f.db)
}

// CreatePhoneRepository constructs mysql phone repository
func (f *RepositoryFabric) CreatePhoneRepository() repositories.PhoneRepository {
	return phoneRepository.New(f.db)
}
