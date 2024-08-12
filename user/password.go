package user

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/tanaka.takuto/go-valueobject-sample/vo"
)

// RawPassword is a raw password
type RawPassword vo.StringValueObject

var (
	uppercaseRegexp   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegexp   = regexp.MustCompile(`[a-z]`)
	digitRegexp       = regexp.MustCompile(`\d`)
	specialCharRegexp = regexp.MustCompile(`[@$!%*?&]`)
	lengthRegexp      = regexp.MustCompile(`^[A-Za-z\d@$!%*?&]{8,20}$`)
)

// NewRawPassword creates a raw password
func NewRawPassword(value string) (*RawPassword, error) {
	if !uppercaseRegexp.MatchString(value) ||
		!lowercaseRegexp.MatchString(value) ||
		!digitRegexp.MatchString(value) ||
		!specialCharRegexp.MatchString(value) ||
		!lengthRegexp.MatchString(value) {
		return nil, errors.New("invalid password")
	}

	rp := RawPassword(vo.NewStringValueObject(value))

	return &rp, nil
}

// Password is a password with salt
type Password vo.HashedWithSaltString

// NewPassword creates a password
func NewPassword(rawPassword RawPassword) Password {
	password := vo.NewHashedStringWithSalt(rawPassword.Value())
	return Password(password)
}

// ValidString checks if the string is valid
func (p Password) ValidString(challengePassword ChallengePassword) error {
	password := vo.HashedWithSaltString(p)
	if err := password.ValidString(challengePassword.Value()); err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

// ChallengePassword is a challenge password
type ChallengePassword vo.StringValueObject

// NewChallengePassword creates a challenge password
func NewChallengePassword(value string) ChallengePassword {
	return ChallengePassword(vo.NewStringValueObject(value))
}
