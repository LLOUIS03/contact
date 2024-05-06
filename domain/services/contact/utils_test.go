package contact

import (
	"testing"

	"github.com/contact/infraestructure/db/repos"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateGetContactByID(t *testing.T) {
	// Test case 1: Valid contact ID
	err := validateContactByID(uuid.New().String())
	assert.NoError(t, err)

	// Test case 2: Invalid contact ID
	err = validateContactByID("invalid-id")
	assert.Error(t, err)
	assert.EqualError(t, err, "Invalid contact ID")
}
func TestValidateContact(t *testing.T) {
	// Test case 1: Valid contact
	contact := repos.Contact{
		Name:     "John",
		Lastname: "Doe",
		Email:    "john.doe@example.com",
		Phone:    "1234567890",
	}
	err := validateContact(contact)
	assert.NoError(t, err)

	// Test case 2: Missing name
	contact = repos.Contact{
		Lastname: "Doe",
		Email:    "john.doe@example.com",
		Phone:    "1234567890",
	}
	err = validateContact(contact)
	assert.Error(t, err)
	assert.EqualError(t, err, "Name is required")

	// Test case 3: Missing lastname
	contact = repos.Contact{
		Name:  "John",
		Email: "john.doe@example.com",
		Phone: "1234567890",
	}
	err = validateContact(contact)
	assert.Error(t, err)
	assert.EqualError(t, err, "Lastname is required")

	// Test case 4: Missing email
	contact = repos.Contact{
		Name:     "John",
		Lastname: "Doe",
		Phone:    "1234567890",
	}
	err = validateContact(contact)
	assert.Error(t, err)
	assert.EqualError(t, err, "Email is required")

	// Test case 5: Missing phone
	contact = repos.Contact{
		Name:     "John",
		Lastname: "Doe",
		Email:    "john.doe@example.com",
	}
	err = validateContact(contact)
	assert.Error(t, err)
	assert.EqualError(t, err, "Phone is required")
}
