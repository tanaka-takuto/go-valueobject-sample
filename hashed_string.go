package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// HashedString is a hashed string
type HashedString string

// NewHashedString creates a hashed string
func NewHashedString(plainStr string) HashedString {
	h := sha512.New()
	_, err := h.Write([]byte(plainStr))
	if err != nil {
		panic(err)
	}

	hashedBytes := h.Sum(nil)
	hashedStr := hex.EncodeToString(hashedBytes)

	return HashedString(fmt.Sprintf("%v:%v", "sha512", hashedStr))
}

// ValidString checks if the string is valid
func (hs HashedString) ValidString(plainStr string) error {
	hashedString := NewHashedString(plainStr)
	if hs != hashedString {
		return fmt.Errorf("invalid hashed string")
	}

	return nil
}

const (
	// saltLength is the length of the salt
	saltLength = 32
)

// HashedWithSaltString is a hashed string with salt
type HashedWithSaltString string

// NewHashedStringWithSalt Create a hashed string with a specified salt
func NewHashedStringWithSalt(plainStr string) HashedWithSaltString {
	salt := newRandomString(saltLength)
	return newHashedStringWithSalt(plainStr, salt)
}

// newHashedStringWithSalt Create a hashed string with a specified salt
func newHashedStringWithSalt(plainStr string, salt string) HashedWithSaltString {
	plainStrWithSalt := plainStr + salt

	hashedString := NewHashedString(plainStrWithSalt)
	return HashedWithSaltString(fmt.Sprintf("%v:%v", hashedString, salt))
}

// ValidString checks if the string is valid
func (hws HashedWithSaltString) ValidString(plainStr string) error {
	splitted := strings.Split(string(hws), ":")
	if len(splitted) != 3 {
		return fmt.Errorf("invalid hashed string")
	}

	algorithm, _, salt := splitted[0], splitted[1], splitted[2]
	if algorithm != "sha512" {
		return fmt.Errorf("invalid hashed string")
	}

	calculatedHash := newHashedStringWithSalt(plainStr, salt)
	if calculatedHash != hws {
		return fmt.Errorf("invalid hashed string")
	}

	return nil
}

// HashedWithPepperSaltString is a hashed string with pepper salt
type HashedWithPepperSaltString string

// NewHashedStringWithPepperSalt Create a hashed string with a specified pepper salt
func NewHashedStringWithPepperSalt(plainStr string) HashedWithPepperSaltString {
	plainWithPepper := plainStr + pepper()
	hashedStringWithPepperSalt := NewHashedStringWithSalt(plainWithPepper)
	return HashedWithPepperSaltString(hashedStringWithPepperSalt)
}

// ValidString checks if the string is valid
func (hwps HashedWithPepperSaltString) ValidString(plainStr string) error {
	plainWithPepper := plainStr + pepper()
	hashedStringWithPepperSalt := HashedWithSaltString(hwps)
	return hashedStringWithPepperSalt.ValidString(plainWithPepper)
}

// pepper returns a pepper string
func pepper() string {
	p := os.Getenv("PEPPER")
	if p == "" {
		panic("PEPPER is not set")
	}
	return p
}
