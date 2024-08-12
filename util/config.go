package util

import "os"

// config is a struct that holds the configuration for the application.
type config struct {
	EncryptKey string
	Pepper     string
}

// c is a global variable that holds the configuration for the application.
var c *config

// GetConfig returns the configuration for the application.
func GetConfig() config {
	if c == nil {
		InitializeConfig()
	}

	return *c
}

// InitializeConfig initializes the configuration by retrieving the encryption key and pepper from environment variables.
func InitializeConfig() {
	encryptKey := os.Getenv("ENCRYPT_KEY")
	pepper := os.Getenv("PEPPER")
	if encryptKey == "" {
		panic("ENCRYPT_KEY is not set")
	}
	if pepper == "" {
		panic("PEPPER is not set")
	}

	c = &config{
		EncryptKey: encryptKey,
		Pepper:     pepper,
	}
}
