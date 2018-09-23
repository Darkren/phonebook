package controllers

import (
	"github.com/Darkren/phonebook/models"
	"github.com/Darkren/phonebook/repositories"
	"github.com/Darkren/phonebook/wsjsonrpc/msg"
)

// UserController is a controller which incapsulates all
// actions connected with users
type UserController struct {
	UserRepo repositories.UserRepository
}

// List returns all the users from the persistent storage
func (c *UserController) List(args *msg.UserListArgs,
	response *msg.UserListResponse) error {
	var err error
	*response, err = c.UserRepo.List()

	return err
}

// Get returns single user by its ID
func (c *UserController) Get(args *msg.GetUserArgs,
	response *msg.GetUserResponse) error {
	var err error
	*response, err = c.UserRepo.Get(args.ID)

	return err
}

// Add adds the user to the persistent storage
func (c *UserController) Add(args *msg.AddUserArgs,
	response *msg.AddUserResponse) error {
	var err error
	var id int64
	user := models.User(*(*args))
	id, err = c.UserRepo.Add(&user)
	idResp := msg.AddUserResponse(id)
	*response = idResp

	return err
}

// Update updates all the fields of the passed user by its ID
func (c *UserController) Update(args *msg.UpdateUserArgs,
	response *msg.UpdateUserResponse) error {
	user := models.User(*(*args))
	err := c.UserRepo.Update(&user)

	result := err == nil
	*response = msg.UpdateUserResponse(result)

	return err
}

// Delete removes user from the persistent storage by its ID
func (c *UserController) Delete(args *msg.DeleteUserArgs,
	response *msg.DeleteUserResponse) error {
	id := int64(*args)

	err := c.UserRepo.Delete(id)

	result := err == nil

	*response = msg.DeleteUserResponse(result)

	return err
}
