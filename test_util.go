package main

import "testing"

// initializeTestConfig initializes the configuration by retrieving the encryption key and pepper from environment variables.
func initializeTestConfig(t *testing.T) {
	t.Setenv("ENCRYPT_KEY", "1234567890abcdefghijklmnopqrstuv")
	t.Setenv("PEPPER", "12345678901234567890123456789012")

	initializeConfig()
}
