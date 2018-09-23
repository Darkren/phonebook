package msg

import "github.com/Darkren/phonebook/models"

// GetPhonesListByUserArgs is a type of args used to
// get phones list by user ID
type GetPhonesListByUserArgs int64

// GetPhonesListByUserResponse is a type of response for the same request
type GetPhonesListByUserResponse []*models.Phone

// AddPhoneArgs is a type of args used to add the phone to storage
type AddPhoneArgs *models.Phone

// AddPhoneResponse is a type of response for the same request
type AddPhoneResponse int64

// UpdatePhoneArgs is a type of args used to update the already stored phone
type UpdatePhoneArgs *models.Phone

// UpdatePhoneResponse is a type of response for the same request
type UpdatePhoneResponse bool

// DeletePhoneArgs is a type of args used to delete the phone from storage
type DeletePhoneArgs int64

// DeletePhoneResponse is a type of response for the same request
type DeletePhoneResponse bool
