package contact

import (
	"net/http"

	"github.com/contact/api/handlers"
	"github.com/contact/domain/services/contact"
	"github.com/labstack/echo/v4"
)

type handler struct {
	contactSvc contact.Contact
}

// @Summary Update
// @Description Update
// @Tags Contact
// @Param id path string true "The id of the contact"
// @Param requestBody body UpdateContact true "Add a contact request"
// @Success 200 {object} nil
// @Router /contact/{id} [put].
func (h *handler) Update(c echo.Context) error {
	req := UpdateContact{}
	reqID := c.Param("id")
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	svcReq := createNewContact(createNewContactReq{
		Name:      req.Name,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
		UpdatedAt: &req.UpdatedAt,
	})

	resp, err := h.contactSvc.UpdateContact(ctx, reqID, svcReq)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Add Contact
// @Description Add a contact
// @Tags Contact
// @Param requestBody body AddContact true "Add a contact request"
// @Success 200 {object} repos.Contact
// @Router /contact [post].
func (h *handler) Add(c echo.Context) error {
	req := AddContact{}
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	svcReq := createNewContact(createNewContactReq{
		Name:     req.Name,
		Lastname: req.Lastname,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	resp, err := h.contactSvc.AddContact(ctx, svcReq)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary GetByID
// @Description Get a contact by id
// @Tags Contact
// @Param id path string true "The id of the contact"
// @Success 200 {object} nil
// @Router /contact/{id} [get].
func (h *handler) GetByID(c echo.Context) error {
	req := c.Param("id")
	ctx := c.Request().Context()

	resp, err := h.contactSvc.GetContactByID(ctx, req)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary GetAll
// @Description Get all the contacts
// @Tags Contact
// @Success 200 {array} repos.Contact
// @Router /contacts [get].
func (h *handler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.contactSvc.GetContacts(ctx)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Delete
// @Description Delete a contact
// @Tags Contact
// @Param id path string true "The id of the contact"
// @Success 200 {object} nil
// @Router /contact/{id} [delete].
func (h *handler) Delete(c echo.Context) error {
	req := c.Param("id")
	ctx := c.Request().Context()

	err := h.contactSvc.DeleteContact(ctx, req)
	if err != nil {
		return handlers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, nil)
}
