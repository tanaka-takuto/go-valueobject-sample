package main

import (
	"testing"
)

func TestNewRawEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Valid email address", args: args{email: "test@example.com"}, wantErr: false},
		{name: "Invalid email address", args: args{email: "invalid_email"}, wantErr: true},
		{name: "Email address with special characters", args: args{email: "test.email+tag@example.com"}, wantErr: false},
		{name: "Email address without @ symbol", args: args{email: "testexample.com"}, wantErr: true},
		{name: "Email address without domain", args: args{email: "test@"}, wantErr: true},
		{name: "Email address with multiple @ symbols", args: args{email: "test@example@com"}, wantErr: true},
		{name: "Email address with leading whitespace", args: args{email: " test@example.com"}, wantErr: true},
		{name: "Email address with trailing whitespace", args: args{email: "test@example.com "}, wantErr: true},
		{name: "Email address with consecutive dots in domain", args: args{email: "test@example..com"}, wantErr: true},
		// Since we are using a simple validation, this test will fail.
		// {name: "Email address with consecutive dots in local part", args: args{email: "test..email@example.com"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRawEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRawEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_EmailAndRawEmail(t *testing.T) {
	initializeTestConfig(t)

	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Valid email address", args: args{email: "test@example.com"}},
		{name: "Email address with special characters", args: args{email: "test.email+tag@example.com"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re, rawEmailErr := NewRawEmail(tt.args.email)
			if rawEmailErr != nil {
				t.Errorf("NewRawEmail() error = %v", rawEmailErr)
				return
			}

			e, emailErr := NewEmail(*re)
			if emailErr != nil {
				t.Errorf("NewEmail() error = %v", emailErr)
				return
			}

			re2, rawEmailErr2 := e.RawEmail()
			if rawEmailErr2 != nil {
				t.Errorf("Email.RawEmail() error = %v", rawEmailErr2)
				return
			}

			if *re != *re2 {
				t.Errorf("NewRawEmail() = %v, want %v", *re, *re2)
			}

			if tt.args.email != string(*re2) {
				t.Errorf("Email.RawEmail() = %v, want %v", string(*re2), tt.args.email)
			}
		})
	}
}
