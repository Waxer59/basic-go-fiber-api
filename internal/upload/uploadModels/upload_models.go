package uploadModels

import (
	"github.com/google/uuid"
)

type Upload struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Ext    string    `json:"-" gorm:"not null"`
	UserID uuid.UUID `json:"-"`
}
