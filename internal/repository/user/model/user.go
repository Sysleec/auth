package model

import "time"

// User is a model for user entity
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Role is a type for user role
type Role int32
