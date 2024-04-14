package models

import "time"

type User struct {
	Email    string    `json:"email"`
	Date     time.Time `json:"date"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Comments string    `json:"comments"`
}
