package models

// Phone is a data model of phone
type Phone struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"userId"`
	Phone  string `json:"phone"`
}
