package contact

import (
	"time"

	"github.com/contact/infraestructure/db/repos"
)

type createNewContactReq struct {
	Name      string
	Lastname  string
	Email     string
	Phone     string
	UpdatedAt *time.Time
}

func createNewContact(req createNewContactReq) repos.Contact {
	contact := repos.Contact{
		Name:     req.Name,
		Lastname: req.Lastname,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	if req.UpdatedAt != nil {
		contact.UpdatedAt = *req.UpdatedAt
	}

	return contact
}
