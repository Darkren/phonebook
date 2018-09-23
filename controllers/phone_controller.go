package controllers

import (
	"github.com/Darkren/phonebook/models"
	"github.com/Darkren/phonebook/repositories"
	"github.com/Darkren/phonebook/wsjsonrpc/msg"
)

// PhoneController is a controller which incapsulates all
// actions connected with phones
type PhoneController struct {
	PhoneRepo repositories.PhoneRepository
}

// ListByUser returns all the phones owned
// by user with the specified ID
func (c *PhoneController) ListByUser(args *msg.GetPhonesListByUserArgs,
	response *msg.GetPhonesListByUserResponse) error {
	userID := int64(*args)

	phones, err := c.PhoneRepo.ListByUser(userID)

	*response = phones

	return err
}

// Add adds the phone to the persistent storage
func (c *PhoneController) Add(args *msg.AddPhoneArgs,
	response *msg.AddPhoneResponse) error {
	phone := models.Phone(*(*args))

	id, err := c.PhoneRepo.Add(&phone)

	*response = msg.AddPhoneResponse(id)

	return err
}

// Update updates all the fields of the passed phone by its ID
func (c *PhoneController) Update(args *msg.UpdatePhoneArgs,
	response *msg.UpdatePhoneResponse) error {
	phone := models.Phone(*(*args))

	err := c.PhoneRepo.Update(&phone)

	result := err == nil

	*response = msg.UpdatePhoneResponse(result)

	return err
}

// Delete removes phone from the persistent storage by its ID
func (c *PhoneController) Delete(args *msg.DeletePhoneArgs,
	response *msg.DeletePhoneResponse) error {
	err := c.PhoneRepo.Delete(int64(*args))

	result := err == nil

	*response = msg.DeletePhoneResponse(result)

	return err
}
