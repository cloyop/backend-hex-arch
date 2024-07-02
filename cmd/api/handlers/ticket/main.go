package ticket

import (
	"backend-hex-arch/pkg/domain"
	"backend-hex-arch/pkg/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ticketHandler struct {
	ticketService ports.TicketService
}

func NewTicketHandler(s ports.TicketService) *ticketHandler {
	return &ticketHandler{
		ticketService: s,
	}
}

func (h *ticketHandler) GetAll(c echo.Context) error {
	// http logic
	tickets, err := h.ticketService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tickets)
}
func (h *ticketHandler) Get(c echo.Context) error {
	// http logic
	t, is, err := h.ticketService.Get(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if !is {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Not exists",
		})
	}
	return c.JSON(http.StatusOK, t)
}
func (h *ticketHandler) Create(c echo.Context) error {
	// HttpLogic
	t := &domain.Ticket{}
	if err := h.ticketService.Create(t); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, t)
}
func (h *ticketHandler) Edit(c echo.Context) error {
	// http logic
	t := &domain.Ticket{}
	if err := h.ticketService.Edit(c.Param("id"), t); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, t)
}
func (h *ticketHandler) Delete(c echo.Context) error {
	// http logic
	if err := h.ticketService.Delete(c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
