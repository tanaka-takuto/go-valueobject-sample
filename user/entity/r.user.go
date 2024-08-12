package user

type UserRepository interface {
	FindUserByID(userID UserID) (*User, error)
	DeleteUser(userID DeletableUserID) error
}
