package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoleCreationScenarios(t *testing.T) {
	tests := []struct {
		name      string
		roleName  string
		expectErr bool
	}{
		{"Correct input", "Admin", false},
		{"Missing name", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			role, err := NewRole(tt.roleName)

			if tt.expectErr {
				assert.NotNil(t, err, "expected an error but got nil")
			} else {
				require.Nil(t, err, "unexpected error creating user")
				assert.Equal(t, tt.roleName, role.Name)
				assert.NotNil(t, role.ID)
			}
		})
	}
}
