package models

import "testing"

func TestCreateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateUserRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			req: CreateUserRequest{
				Username: "testuser",
				Email:    "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "Missing username",
			req: CreateUserRequest{
				Email: "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "Missing email",
			req: CreateUserRequest{
				Username: "testuser",
			},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     CreateUserRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     UpdateUserRequest
		wantErr bool
	}{
		{
			name: "Valid with username",
			req: UpdateUserRequest{
				Username: "newusername",
			},
			wantErr: false,
		},
		{
			name: "Valid with email",
			req: UpdateUserRequest{
				Email: "new@example.com",
			},
			wantErr: false,
		},
		{
			name: "Valid with both",
			req: UpdateUserRequest{
				Username: "newusername",
				Email:    "new@example.com",
			},
			wantErr: false,
		},
		{
			name:    "Empty request",
			req:     UpdateUserRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
