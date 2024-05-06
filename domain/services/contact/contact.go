package contact

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/contact/domain/domainerrors"
	"github.com/contact/infraestructure/db/repos"
	"github.com/google/uuid"
)

// Contact is the interface that wraps the transaction methods
type Contact interface {
	GetContacts(context.Context) ([]repos.Contact, error)
	GetContactByID(context.Context, string) (*repos.Contact, error)
	AddContact(context.Context, repos.Contact) (*repos.Contact, error)
	UpdateContact(context.Context, string, repos.Contact) (*repos.Contact, error)
	DeleteContact(context.Context, string) error
}

type contact struct {
	queries repos.Querier
}

// NewService creates a new contact service
func NewService(queries repos.Querier) Contact {
	return &contact{
		queries: queries,
	}
}

// GetContactByID returns a contact by its ID
func (c *contact) GetContactByID(ctx context.Context, contactID string) (*repos.Contact, error) {
	if err := validateContactByID(contactID); err != nil {
		return nil, err
	}

	contactUUID := uuid.MustParse(contactID)

	resp, err := c.queries.GetContactByID(ctx, contactUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, domainerrors.NewInternalServerError(err)
	}

	return &resp, nil
}

// AddContact adds a new contact
func (c *contact) AddContact(ctx context.Context, model repos.Contact) (*repos.Contact, error) {
	fmt.Println("model: ", model)
	if err := validateContact(model); err != nil {
		return nil, err
	}

	newID := uuid.New()

	resp, err := c.queries.CreateContact(ctx, repos.CreateContactParams{
		ID:       newID,
		Name:     model.Name,
		Lastname: model.Lastname,
		Email:    model.Email,
		Phone:    model.Phone,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainerrors.NewNoDataError(err)
		}
		return nil, domainerrors.NewInternalServerError(err)
	}

	return &resp, nil
}

// UpdateContact updates a contact
func (c *contact) UpdateContact(ctx context.Context, id string, model repos.Contact) (*repos.Contact, error) {
	if err := validateContact(model); err != nil {
		return nil, err
	}

	if err := validateContactByID(id); err != nil {
		return nil, err
	}

	contactUUID := uuid.MustParse(id)

	updContact := repos.UpdateContactParams{
		ID:        contactUUID,
		Name:      model.Name,
		Lastname:  model.Lastname,
		Phone:     model.Phone,
		UpdatedAt: model.UpdatedAt,
	}

	resp, err := c.queries.UpdateContact(ctx, updContact)
	if err != nil {
		return nil, domainerrors.NewInternalServerError(err)
	}

	return &resp, nil
}

// DeleteContact deletes a contact
func (c *contact) DeleteContact(ctx context.Context, contactID string) error {
	if err := validateContactByID(contactID); err != nil {
		return err
	}

	contactUUID := uuid.MustParse(contactID)

	err := c.queries.DeleteContact(ctx, contactUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domainerrors.NewNoDataError(err)
		}
		return err
	}

	return nil
}

// GetContacts returns all contacts
func (c *contact) GetContacts(ctx context.Context) ([]repos.Contact, error) {
	resp, err := c.queries.GetContacts(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domainerrors.NewNoDataError(err)
		}
		return nil, err
	}
	return resp, nil
}
