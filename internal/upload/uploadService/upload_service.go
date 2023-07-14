package uploadService

import (
	"errors"

	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadModels"
)

func CreateUpload(userID uuid.UUID, fileName uuid.UUID, ext string) (*uploadModels.Upload, error) {
	db := database.DB

	upload := uploadModels.Upload{
		ID:     fileName,
		Ext:    ext,
		UserID: userID,
	}

	err := db.Create(&upload).Error

	if err != nil {
		return nil, errors.New("Invalid upload")
	}

	return &upload, nil
}
