package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserCreationScenarios(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		email     string
		password  string
		expectErr bool
	}{
		{"Valid user", "John Doe", "doe@email.com", "12345678", false},
		{"Missing name", "", "doe@email.com", "12345678", true},
		{"Invalid email", "John Doe", "invalid-email", "12345678", true},
		{"Short password", "John Doe", "doe@email.com", "123", true},
		{"Missing created_at", "John Doe", "doe@email.com", "123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roleId, err := uuid.NewV7()
			require.Nil(t, err, "failed to generate UUID")

			user, err := NewUser(tt.username, tt.email, tt.password, roleId)

			if tt.expectErr {
				assert.NotNil(t, err, "expected an error but got nil")
			} else {
				require.Nil(t, err, "unexpected error creating user")
				assert.Equal(t, tt.username, user.Name)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.password, user.Password)
				assert.Equal(t, roleId, user.RoleID)
				assert.NotEmpty(t, user.CreatedAt)
				assert.NotEmpty(t, user.UpdatedAt)
			}
		})
	}
}
