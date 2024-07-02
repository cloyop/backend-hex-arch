package ticket

import (
	"backend-hex-arch/pkg/domain"
	"backend-hex-arch/pkg/ports"
)

type service struct {
	repository ports.TicketRepository
}

func NewPlayerService(repository ports.TicketRepository) ports.TicketService {
	return &service{
		repository: repository,
	}
}
func (s *service) GetAll() (*[]domain.Ticket, error) {
	// bussiness Logic
	return s.repository.GetAll()
}
func (s *service) Get(id string) (*domain.Ticket, bool, error) {
	// bussiness Logic
	_, _ = s.repository.Get(id)
	return nil, false, nil
}
func (s *service) Create(t *domain.Ticket) error {
	// bussiness Logic
	t.GenID()
	return s.repository.Insert(t)
}
func (s *service) Edit(id string, i interface{}) error {
	// bussiness Logic
	return s.repository.Edit(id, i)
}
func (s *service) Delete(id string) error {
	// bussiness Logic
	return s.repository.Delete(id)
}
