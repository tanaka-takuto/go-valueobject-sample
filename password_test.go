package main

import "testing"

func TestPassword_ValidString(t *testing.T) {
	type args struct {
		rawPassword       string
		challengePassword string
	}
	type wants struct {
		rawPasswordError       bool
		challengePasswordError bool
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{name: "Valid password", args: args{rawPassword: "Abc123!@", challengePassword: "Abc123!@"}, wants: wants{rawPasswordError: false, challengePasswordError: false}},
		{name: "Invalid challenge password", args: args{rawPassword: "Abc123!@", challengePassword: "wrong_challenge"}, wants: wants{rawPasswordError: false, challengePasswordError: true}},
		{name: "Invalid raw password: insufficient length", args: args{rawPassword: "Abc123!", challengePassword: "Abc123!"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: excessive length", args: args{rawPassword: "Abc123!Abc123!Abc123!Abc123!Abc123!", challengePassword: "Abc123!Abc123!Abc123!Abc123!Abc123!"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: missing uppercase letter", args: args{rawPassword: "abc123!@", challengePassword: "abc123!@"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: missing lowercase letter", args: args{rawPassword: "ABC123!@", challengePassword: "ABC123!@"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: missing digit", args: args{rawPassword: "Abcdefg!@", challengePassword: "Abcdefg!@"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: missing symbol", args: args{rawPassword: "Abc123456", challengePassword: "Abc123456"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
		{name: "Invalid raw password: invalid character", args: args{rawPassword: "Abc123!@#", challengePassword: "Abc123!@#"}, wants: wants{rawPasswordError: true, challengePasswordError: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawPassword, err := NewRawPassword(tt.args.rawPassword)
			if (err != nil) != tt.wants.rawPasswordError {
				t.Errorf("Password.ValidString() error = %v, rawPasswordError %v", err, tt.wants.rawPasswordError)
			}
			if err != nil {
				return
			}
			password := NewPassword(rawPassword)

			challengePassword := NewChallengePassword(tt.args.challengePassword)
			if err := password.ValidString(challengePassword); (err != nil) != tt.wants.challengePasswordError {
				t.Errorf("Password.ValidString() error = %v, challengePasswordError %v", err, tt.wants.challengePasswordError)
			}
		})
	}
}
