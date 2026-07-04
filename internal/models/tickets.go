package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tickets struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"default:open" json:"status"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tickets) BeforeCreate(tx *gorm.DB) error{
	if(t.ID == uuid.Nil){
		t.ID = uuid.New()
	}
	return nil
}
