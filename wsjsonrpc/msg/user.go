package msg

import "github.com/Darkren/phonebook/models"

// UserListArgs is a type of args used to get the users list
type UserListArgs struct{}

// UserListResponse is a type of response for the same request
type UserListResponse []*models.User

// GetUserArgs is a type of args used to get the user by its ID
type GetUserArgs struct {
	ID int64 `json:"id"`
}

// GetUserResponse is a type of response for the same request
type GetUserResponse *models.User

// AddUserArgs is a type of args used to add the user to storage
type AddUserArgs *models.User

// AddUserResponse is a type of response for the same request
type AddUserResponse int64

// UpdateUserArgs is a type of args used to update the already stored user
type UpdateUserArgs *models.User

// UpdateUserResponse is a type of response for the same request
type UpdateUserResponse bool

// DeleteUserArgs is a type of args used to delete the user by its ID
type DeleteUserArgs int64

// DeleteUserResponse is a type of response for the same request
type DeleteUserResponse bool
