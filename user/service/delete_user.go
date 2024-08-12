package user_service

import user "github.com/tanaka.takuto/go-valueobject-sample/user/entity"

type DeleteUserInput struct {
	UserID user.UserID
	Actor  user.User
}

type DeleteUserOutput struct{}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(input DeleteUserInput) (*DeleteUserOutput, error) {
	u, err := s.userRepository.FindUserByID(input.UserID)
	if err != nil {
		return nil, err
	}

	deletableUserID, err := u.Delete(input.Actor)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.DeleteUser(*deletableUserID)
	if err != nil {
		return nil, err
	}

	return &DeleteUserOutput{}, nil
}
