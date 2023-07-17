package uploadService

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadModels"
)

func DeleteUpload(uploadID string, userID uuid.UUID) (*uploadModels.Upload, error) {
	db := database.DB

	upload, err := GetUpload(uploadID, userID.String())

	if err != nil {
		return nil, fiber.ErrNotFound
	}

	err = db.Delete(&upload, "id = ?", uploadID).Error

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(dir + "/uploads/" + upload.ID.String() + "." + upload.Ext)

	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return nil, nil
}

func GetUpload(uploadID string, userID string) (*uploadModels.Upload, error) {
	db := database.DB
	var upload uploadModels.Upload

	err := db.Model(&uploadModels.Upload{}).Where("id = ? AND user_id = ?", uploadID, userID).First(&upload).Error

	return &upload, err
}

func GetAllUploads(userID string) ([]uploadModels.Upload, error) {
	db := database.DB
	var uploads []uploadModels.Upload
	err := db.Model(&uploadModels.Upload{}).Where("user_id = ?", userID).Find(&uploads).Error
	return uploads, err
}

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
