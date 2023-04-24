package uploadModels

import (
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
)

type Upload struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key;"`
	File_name uuid.UUID       `json:"file_name" gorm:"not null"`
	UserID    uuid.UUID       `json:"userId" `
	User      userModels.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}
