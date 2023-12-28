package main

import "testing"

func TestHashedString_ValidString(t *testing.T) {
	type args struct {
		plainStr     string
		challengeStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Valid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxyz"}, wantErr: false},
		{name: "Invalid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxy"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := NewHashedString(tt.args.plainStr)
			if err := hs.ValidString(tt.args.challengeStr); (err != nil) != tt.wantErr {
				t.Errorf("HashedString.ValidString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashedWithSaltString_ValidString(t *testing.T) {
	type args struct {
		plainStr     string
		challengeStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Valid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxyz"}, wantErr: false},
		{name: "Invalid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxy"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hws := NewHashedStringWithSalt(tt.args.plainStr)
			if err := hws.ValidString(tt.args.challengeStr); (err != nil) != tt.wantErr {
				t.Errorf("HashedWithSaltString.ValidString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashedWithPepperSaltString_ValidString(t *testing.T) {
	t.Setenv("PEPPER", "12345678901234567890123456789012")

	type args struct {
		plainStr     string
		challengeStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Valid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxyz"}, wantErr: false},
		{name: "Invalid hashed string", args: args{plainStr: "abcdefghijklmnopqrstuvwxyz", challengeStr: "abcdefghijklmnopqrstuvwxy"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hwps := NewHashedStringWithPepperSalt(tt.args.plainStr)
			if err := hwps.ValidString(tt.args.challengeStr); (err != nil) != tt.wantErr {
				t.Errorf("HashedStringWithPepperSalt.ValidString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
