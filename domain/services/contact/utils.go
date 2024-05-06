package contact

import (
	"github.com/contact/domain/domainerrors"
	"github.com/contact/infraestructure/db/repos"
	"github.com/google/uuid"
)

func validateContactByID(contactID string) error {
	domainErr := domainerrors.NewValidationError()
	if _, err := uuid.Parse(contactID); err != nil {
		domainErr.AddErrorMsg("Invalid contact ID")
	}

	if domainErr.IsSuccess() {
		return nil
	}

	return domainErr
}

func validateContact(contact repos.Contact) error {
	domainErr := domainerrors.NewValidationError()
	if contact.Name == "" {
		domainErr.AddErrorMsg("Name is required")
	}

	if contact.Lastname == "" {
		domainErr.AddErrorMsg("Lastname is required")
	}

	if contact.Email == "" {
		domainErr.AddErrorMsg("Email is required")
	}

	if contact.Phone == "" {
		domainErr.AddErrorMsg("Phone is required")
	}

	if domainErr.IsSuccess() {
		return nil
	}

	return domainErr
}
