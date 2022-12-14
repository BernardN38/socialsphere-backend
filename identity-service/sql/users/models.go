// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package users

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type UserProfileImage struct {
	UserID  uuid.UUID
	ImageID uuid.UUID
}
