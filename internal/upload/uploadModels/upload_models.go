package uploadModels

import (
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
)

type Upload struct {
	ID     uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Ext    string          `json:"-" gorm:"not null"`
	UserID uuid.UUID       `json:"-"`
	User   userModels.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE" json:"-"`
}
