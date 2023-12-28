package main

import (
	"errors"
	"fmt"
	"regexp"
)

// RawPassword is a raw password
type RawPassword string

var (
	uppercaseRegexp   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegexp   = regexp.MustCompile(`[a-z]`)
	digitRegexp       = regexp.MustCompile(`\d`)
	specialCharRegexp = regexp.MustCompile(`[@$!%*?&]`)
	lengthRegexp      = regexp.MustCompile(`^[A-Za-z\d@$!%*?&]{8,20}$`)
)

// NewRawPassword creates a raw password
func NewRawPassword(value string) (RawPassword, error) {
	if !uppercaseRegexp.MatchString(value) ||
		!lowercaseRegexp.MatchString(value) ||
		!digitRegexp.MatchString(value) ||
		!specialCharRegexp.MatchString(value) ||
		!lengthRegexp.MatchString(value) {
		return "", errors.New("invalid password")
	}

	return RawPassword(value), nil
}

// Password is a password with salt
type Password HashedWithSaltString

// NewPassword creates a password
func NewPassword(rawPassword RawPassword) Password {
	password := NewHashedStringWithSalt(string(rawPassword))
	return Password(password)
}

// ValidString checks if the string is valid
func (p Password) ValidString(challengePassword ChallengePassword) error {
	password := HashedWithSaltString(p)
	if err := password.ValidString(string(challengePassword)); err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

// ChallengePassword is a challenge password
type ChallengePassword string

// NewChallengePassword creates a challenge password
func NewChallengePassword(value string) ChallengePassword {
	return ChallengePassword(value)
}
