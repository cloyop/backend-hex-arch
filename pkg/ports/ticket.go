package ports

import (
	"backend-hex-arch/pkg/domain"
)

type TicketService interface {
	Get(string) (*domain.Ticket, bool, error)
	GetAll() (*[]domain.Ticket, error)
	Create(*domain.Ticket) error
	Edit(string, interface{}) error
	Delete(string) error
}
type TicketRepository interface {
	Get(string) (*domain.Ticket, error)
	GetAll() (*[]domain.Ticket, error)
	Insert(*domain.Ticket) error
	Edit(string, interface{}) error
	Delete(string) error
}
