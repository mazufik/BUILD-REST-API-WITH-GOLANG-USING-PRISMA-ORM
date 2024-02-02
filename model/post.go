package model

import "time"

type Post struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Published   bool
	Description string
}
