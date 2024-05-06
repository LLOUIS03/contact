package api

import (
	"github.com/contact/domain/services/contact"
	"github.com/labstack/echo/v4"
)

type API struct {
	server   *echo.Echo
	services Services
}

// Services is a struct that contains all the services
type Services struct {
	ContactSvc contact.Contact
}
