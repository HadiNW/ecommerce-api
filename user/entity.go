package user

import "time"

// User ...
type User struct {
	ID        int
	Username  string
	Password  string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
