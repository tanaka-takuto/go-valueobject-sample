package user

import "github.com/tanaka.takuto/go-valueobject-sample/vo"

// UserID is a user ID
type UserID vo.IntValueObject

// NewUserID creates a user ID
func NewUserID(value int) UserID {
	return UserID(vo.NewIntValueObject(value))
}

// DeletableUserID is a user ID that can be deleted
type DeletableUserID struct {
	UserID
}
