package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Address(t *testing.T) {
	id, _ := uuid.NewV7()
	address, err := NewAddress(id, "street", "city", "st", "22222222")
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, address.UserId, id)
	assert.Equal(t, address.Street, "street")
	assert.Equal(t, address.City, "city")
	assert.Equal(t, address.State, "st")
	assert.Equal(t, address.ZipCode, "22222222")
	assert.NotNil(t, address.Id)
}

func Test_Address_Failure(t *testing.T) {
	t.Run("it should fail if a required field is missing", func(t *testing.T) {
		validId, _ := uuid.NewV7()
		tests := []Address{
			{uuid.Nil, uuid.Nil, "street", "city", "st", "22222222"},
			{uuid.Nil, validId, "", "city", "st", "22222222"},
			{uuid.Nil, validId, "street", "", "st", "22222222"},
			{uuid.Nil, validId, "street", "city", "", "22222222"},
			{uuid.Nil, validId, "street", "city", "st", ""},
		}

		for _, test := range tests {
			address, err := NewAddress(test.UserId, test.Street, test.City, test.State, test.ZipCode)
			assert.NotNil(t, err)
			assert.Empty(t, address)
		}
	})
}
