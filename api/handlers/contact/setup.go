package contact

import (
	"github.com/contact/domain/services/contact"
	"github.com/labstack/echo/v4"
)

const (
	getContactById = "/contact/:id"
	updateRoute    = "/contact/:id"
	deleteRoute    = "/contact/:id"
	getContacts    = "/contacts"
	addRoute       = "/contact"
)

func SetupHandler(e *echo.Echo, svc contact.Contact) {
	hdl := handler{contactSvc: svc}
	e.GET(getContactById, hdl.GetByID)
	e.GET(getContacts, hdl.GetAll)
	e.DELETE(deleteRoute, hdl.Delete)
	e.POST(addRoute, hdl.Add)
	e.PUT(updateRoute, hdl.Update)
}
