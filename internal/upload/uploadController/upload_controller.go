package uploadController

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/google/uuid"
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
	upload.Get("/:id", getFileUpload)
	upload.Get("/", getFileUploads)
}

// Delete a file
//
//	@Description	Delete a file
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//	@Param			id				path	string	true	"File ID"
//
//	@Param			Authorization	header	string	true	"Bearer {token}"
//
//	@Router			/upload/:id [delete]
func deleteUpload(c *fiber.Ctx) error {
	userJWT, err := jwtUtils.GetAndParseJwt(c)

	if err != nil {
		return err
	}

	uploadDelete, err := uploadService.DeleteUpload(c.Params("id"), userJWT.ID)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully", "data": uploadDelete})
}

// Get a file
//
//	@Description	Get a file
//	@Tags			Upload
//	@Accept			json
//	@Produce		mpfd
//	@Param			id				path	string	true	"File ID"
//
//	@Param			Authorization	header	string	true	"Bearer {token}"
//
//	@Router			/upload/:id [get]
func getFileUpload(c *fiber.Ctx) error {
	userJWT, err := jwtUtils.GetAndParseJwt(c)

	if err != nil {
		return err
	}

	upload, err := uploadService.GetUpload(c.Params("id"), userJWT.ID.String())

	if err != nil {
		return err
	}

	return c.SendFile(fmt.Sprintf("./uploads/%s.%s", upload.ID, upload.Ext))
}

// Get all files
//
//	@Description	Get all files
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//
//	@Param			Authorization	header	string	true	"Bearer {token}"
//
//	@Router			/upload [get]
func getFileUploads(c *fiber.Ctx) error {
	userJWT, err := jwtUtils.GetAndParseJwt(c)

	if err != nil {
		return err
	}

	uploads, err := uploadService.GetAllUploads(userJWT.ID.String())

	if err != nil {
		return err
	}

	ids := make([]string, len(uploads))
	for i, d := range uploads {
		ids[i] = d.ID.String()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully", "data": ids})
}

// Upload a file
//
//	@Description	Upload a file
//	@Tags			Upload
//	@Accept			mpfd
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Router			/upload/ [post]
func uploadFile(c *fiber.Ctx) error {
	c.Accepts("multipart/form-data")

	file, err := c.FormFile(FORM_PARAM_NAME)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userJWT, err := jwtUtils.GetAndParseJwt(c)

	if err != nil {
		return err
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
