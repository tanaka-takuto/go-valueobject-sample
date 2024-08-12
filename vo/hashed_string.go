package vo

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tanaka.takuto/go-valueobject-sample/util"
)

// HashedString is a hashed string
type HashedString StringValueObject

// NewHashedString creates a hashed string
func NewHashedString(plainStr string) HashedString {
	h := sha512.New()
	_, err := h.Write([]byte(plainStr))
	if err != nil {
		panic(err)
	}

	hashedBytes := h.Sum(nil)
	hashedStr := hex.EncodeToString(hashedBytes)

	hashedStrWithAlgorithm := fmt.Sprintf("%v:%v", "sha512", hashedStr)

	return HashedString(NewStringValueObject(hashedStrWithAlgorithm))
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
type HashedWithSaltString StringValueObject

// NewHashedStringWithSalt Create a hashed string with a specified salt
func NewHashedStringWithSalt(plainStr string) HashedWithSaltString {
	salt := util.NewRandomString(saltLength)
	return newHashedStringWithSalt(plainStr, salt)
}

// newHashedStringWithSalt Create a hashed string with a specified salt
func newHashedStringWithSalt(plainStr string, salt string) HashedWithSaltString {
	plainStrWithSalt := plainStr + salt

	hashedString := NewHashedString(plainStrWithSalt)

	hashedStringWithSalt := fmt.Sprintf("%v:%v", hashedString.value, salt)

	return HashedWithSaltString(NewStringValueObject(hashedStringWithSalt))
}

// ValidString checks if the string is valid
func (hws HashedWithSaltString) ValidString(plainStr string) error {
	splitted := strings.Split(string(hws.value), ":")
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
type HashedWithPepperSaltString StringValueObject

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
	p := util.GetConfig().Pepper
	if p == "" {
		panic("PEPPER is not set")
	}
	return p
}
