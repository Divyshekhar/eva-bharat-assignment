package services

import (
	"errors"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/dto"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/models"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/repository"
	"github.com/google/uuid"
)

type TicketService interface {
	Create(userID uuid.UUID, req dto.CreateTicketRequest) error
	GetAll(userID uuid.UUID) ([]models.Tickets, error)
	GetById(userID, ticketID uuid.UUID) (*models.Tickets, error)
	UpdateStatus(userID, ticketID uuid.UUID, status string) error
}

type ticketService struct {
	repo repository.TicketRepository
}

func (s *ticketService) Create(userID uuid.UUID, req dto.CreateTicketRequest) error {
	ticket := &models.Tickets{
		Title:       req.Title,
		Description: req.Description,
		Status:      "open",
		UserID:      userID,
	}
	return s.repo.Create(ticket)

}
func (s *ticketService) GetAll(userID uuid.UUID) ([]models.Tickets, error) {
	return s.repo.GetByUserID(userID)
}
func (s *ticketService) GetById(userID, ticketID uuid.UUID) (*models.Tickets, error) {
	ticket, err := s.repo.GetByID(ticketID)
	if err != nil {
		return nil, err
	}
	if ticket.UserID != userID {
		return nil, errors.New("Unauthorized")
	}
	return ticket, nil
}
func (s *ticketService) UpdateStatus(userID, ticketID uuid.UUID, status string) error {
	ticket, err := s.repo.GetByID(ticketID)
	if err != nil {
		return err
	}
	if ticket.UserID != userID {
		return errors.New("Unauthorized")
	}
	switch ticket.Status {
	case "open":
		if status != "in_progress" {
			return errors.New("Invalid status transition")
		}
	case "in_progress":
		if status != "closed" {
			return errors.New("Invalid status transition")
		}
	case "closed":
		return errors.New("Closed ticket cannot be reopened")
	}
	ticket.Status = status
	return s.repo.Update(ticket)
}

func NewTicketService(ticketRepository repository.TicketRepository) TicketService {
	return &ticketService{repo: ticketRepository}
}
