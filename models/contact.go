package models

import "time"

type Contact struct {
	Id          int
	Name        string
	PhoneNumber string
	Slug        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (contact *Contact) TableName() string {
	return "contacts"
}
