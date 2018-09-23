package fabric

import "github.com/Darkren/phonebook/repositories"

// RepositoryFabric is an abstract factory of repositories
type RepositoryFabric interface {
	CreateUserRepository() repositories.UserRepository
	CreatePhoneRepository() repositories.PhoneRepository
}
