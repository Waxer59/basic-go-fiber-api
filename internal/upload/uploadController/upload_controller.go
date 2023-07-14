package uploadController

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/internal/helpers"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadModels"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadService"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/jwtUtils"
	"github.com/waxer59/basic-go-fiber-api/router/routeHandlers"
	"github.com/waxer59/basic-go-fiber-api/router/routeMiddlewares"
)

const FORM_PARAM_NAME = "file"

func Setup(router fiber.Router) {
	upload := router.Group("/upload", skip.New(routeHandlers.ProtectRouteHandler, routeMiddlewares.ProtectRouteMiddleware))

	upload.Post("/", uploadFile)
	upload.Delete("/:id", deleteUpload)
}

func deleteUpload(c *fiber.Ctx) error {
	db := database.DB
	var upload uploadModels.Upload

	token, err := helpers.GetJwtToken(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userJWT, err := jwtUtils.ParseJwt(token)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = db.Model(&uploadModels.Upload{}).Preload("User").Where("id = ?", c.Params("id")).First(&upload).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if upload.User.ID != userJWT.ID {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = db.Delete(&upload, "id = ?", c.Params("id")).Error

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(dir + "/uploads/" + upload.ID.String() + "." + upload.Ext)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully", "data": upload})
}

func uploadFile(c *fiber.Ctx) error {
	c.Accepts("multipart/form-data")

	file, err := c.FormFile(FORM_PARAM_NAME)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := helpers.GetJwtToken(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	userJWT, err := jwtUtils.ParseJwt(token)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	fileName := uuid.New()
	ext := strings.Split(file.Filename, ".")[1]

	data, err := uploadService.CreateUpload(userJWT.ID, fileName, ext)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s.%s", fileName, ext)); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data":    data,
	})
}
