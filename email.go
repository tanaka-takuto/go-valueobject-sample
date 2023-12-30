package main

import (
	"errors"
	"regexp"
)

var (
	// emailRegex is a regular expression that matches the email format.
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)
)

// isEmailValid checks if the email is valid
func isEmailValid(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email address")
	}

	return nil
}

// RawEmail is a type that holds the raw email address.
type RawEmail string

// NewRawEmail returns a new RawEmail.
func NewRawEmail(email string) (*RawEmail, error) {
	if err := isEmailValid(email); err != nil {
		return nil, err
	}

	re := RawEmail(email)
	return &re, nil
}

// Email is a type that holds the encrypted email address.
type Email EncryptedBytes

// NewEmail returns a new Email.
func NewEmail(rawEmail RawEmail) (*Email, error) {
	ds := NewDecryptedString(string(rawEmail))
	eb, err := ds.Encrypt()
	if err != nil {
		return nil, err
	}

	return (*Email)(eb), nil
}

// Decrypt returns the decrypted email address.
func (e *Email) RawEmail() (*RawEmail, error) {
	eb := EncryptedBytes(*e)
	ds, err := eb.Decrypt()
	if err != nil {
		return nil, err
	}

	return NewRawEmail(string(*ds))
}
