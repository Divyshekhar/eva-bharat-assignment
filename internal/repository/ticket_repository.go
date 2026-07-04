package repository

import (
	"github.com/Divyshekhar/eva-bharat-assignment/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *models.Tickets) error
	GetByID(id uuid.UUID) (*models.Tickets, error)
	GetByUserID(userID uuid.UUID) ([]models.Tickets, error)
	Update(ticket *models.Tickets) error
}

type ticketRepository struct {
	db *gorm.DB
}

func (r *ticketRepository) Create(ticket *models.Tickets) error {
	return r.db.Create(ticket).Error
}
func (r *ticketRepository) GetByID(id uuid.UUID) (*models.Tickets, error) {
	var ticket models.Tickets
	err := r.db.First(&ticket, "id = ?", id).Error
	if err != nil{
		return nil, err
	}
	return &ticket, nil
}
func (r *ticketRepository) GetByUserID(userId uuid.UUID) ([]models.Tickets, error) {
	var tickets []models.Tickets
	err := r.db.Where("user_id = ?", userId).Find(&tickets).Error
	if err != nil{
		return nil, err
	}
	return tickets, nil
}
func (r *ticketRepository) Update(ticket *models.Tickets) error {
	return r.db.Save(ticket).Error
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}
