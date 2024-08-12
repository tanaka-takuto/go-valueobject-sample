package user_service

import user "github.com/tanaka.takuto/go-valueobject-sample/user/entity"

type UserService struct {
	userRepository user.UserRepository
}
