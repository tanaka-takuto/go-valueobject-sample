package main

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	initializeTestConfig(t)

	tests := []struct {
		name     string
		plainStr string
	}{
		{name: "Testing with English word", plainStr: "Hello"},
		{name: "Testing with Programming term", plainStr: "Programming"},
		{name: "Testing with Encryption term", plainStr: "Encryption"},
		{name: "Testing with Japanese greeting", plainStr: "こんにちは"},
		{name: "Testing with Chinese greeting", plainStr: "你好"},
		{name: "Testing with email format", plainStr: "test@example.com"},
		{name: "Testing with special characters", plainStr: "!@#$%^&*()"},
		{name: "Testing with empty string", plainStr: ""},
		{name: "Testing with whitespace", plainStr: "   "},
		{name: "Testing with numbers", plainStr: "1234567890"},
		{name: "Testing with long string", plainStr: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, urna id ultricies ultrices, nunc nunc ultricies elit, nec aliquet nunc nunc eu nunc."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := NewDecryptedString(tt.plainStr)
			es, eErr := ds.Encrypt()
			if eErr != nil {
				t.Errorf("Encryption error: %v", eErr)
				return
			}
			got, dErr := es.Decrypt()
			if dErr != nil {
				t.Errorf("Decryption error: %v", dErr)
				return
			}

			if tt.plainStr != string(*got) {
				t.Errorf("Encrypt And Decrypt = %v, want %v", *got, tt.plainStr)
			}
		})
	}
}
