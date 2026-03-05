package controllers

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Upload(c *fiber.Ctx) error {

	var name string
	var url string

	filesData, err := c.MultipartForm()
	if err !=  nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if filesData == nil || len(filesData.File) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No files uploaded"})
	}

	for _, files := range filesData.File {
		for _, file := range files {

			// 1. limit ukuran 5MB
			if file.Size > 5 * 1024 * 1024 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "File size exceeds 5MB"})
			}

			// 2. limit jenis file
			allowedTypes := []string{"image/jpeg", "image/png", "image/jpg"}
			isAllowed := slices.Contains(allowedTypes, file.Header.Get("Content-Type")) 
			
			if !isAllowed {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "File type not allowed"})
			}

			// 3. simpan File 

			splitDots := strings.Split(file.Filename, ".")
			ext := splitDots[len(splitDots) - 1]

			newFileName := uuid.New().String() + "." + ext

			if err := c.SaveFile(file,"./uploads/" + newFileName) ; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
			}

			name = newFileName
			url = "/uploads/" + newFileName

		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"fileName": name,
			"url": url,
		})
}