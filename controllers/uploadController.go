package controllers

import (
	"backend-rsmata-360/usecases"

	"github.com/gofiber/fiber/v2"
)

// func UploadSingle(c *fiber.Ctx) error{
// 	var name string
// 	var url string

// 	filesData, err := c.MultipartForm()

// 	if err != nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":"failed",
// 			"message":err.Error(),
// 		})
// 	}

// 	if filesData == nil || len(filesData.File) == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":"failed",
// 			"message":"No files uploaded",
// 		})
// 	}

// 	for _,files := range filesData.File{
// 		for _, file := range files{

// 			fileMeta := usecases.FileMeta{
// 				Filename: file.Filename,
// 				Size: file.Size,
// 				ContentType: file.Header.Get("Content-Type"),
// 			}

// 			uploadResult, err := usecases.UploadCase(fileMeta)

// 			if err != nil{
// 				if err.Error() == "file size exceeds 5mb"{
// 				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 					"status":"failed",
// 					"message":"file size exceeds 5mb",
// 				})
// 			}

// 			if err.Error() == "file type not allowed"{
// 				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 					"status":"failed",
// 					"message":"file type not allowed",
// 				})
// 			}
// 			}

// 			if errr := c.SaveFile(file, "./uploads/" + uploadResult.Filename); errr != nil{ return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"status":"failed",
// 				"message":errr.Error(),
// 			})}

// 			name = uploadResult.Filename
// 			url = uploadResult.Url
// 		}
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"filename": name,
// 		"url":url,
// 	})
// }

func Upload(c *fiber.Ctx) error{
	
	filesData, err := c.MultipartForm()

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}

	if filesData == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"No files uploaded",
		})
	}

	files := filesData.File["image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"No file Uploaded",
		})
	}

	file := files[0]

	fileMeta := usecases.FileMeta{
		Filename: file.Filename,
		Size: file.Size,
		ContentType: file.Header.Get("Content-Type"),
	}

	result, err := usecases.UploadCase(fileMeta)

	if err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message": err.Error(),
			})
		}
	
	if err := c.SaveFile(file, "./uploads/" + result.Filename); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"fileName": result.Filename,
		"url": result.Url,
	})
}

// func UploadEXP(c *fiber.Ctx) error {
	
// 	var name string
// 	var url string

// 	filesData, err := c.MultipartForm()
// 	if err !=  nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	if filesData == nil || len(filesData.File) == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No files uploaded"})
// 	}

// 	for _, files := range filesData.File {
// 		for _, file := range files {

// 			// 1. limit ukuran 5MB
// 			if file.Size > 5 * 1024 * 1024 {
// 				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "File size exceeds 5MB"})
// 			}

// 			// 2. limit jenis file
// 			allowedTypes := []string{"image/jpeg", "image/png", "image/jpg"}
// 			isAllowed := slices.Contains(allowedTypes, file.Header.Get("Content-Type")) 
			
// 			if !isAllowed {
// 				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "File type not allowed"})
// 			}

// 			// 3. simpan File 

// 			splitDots := strings.Split(file.Filename, ".")
// 			ext := splitDots[len(splitDots) - 1]

// 			newFileName := uuid.New().String() + "." + ext

// 			if err := c.SaveFile(file,"./uploads/" + newFileName) ; err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 			}

// 			name = newFileName
// 			url = "/uploads/" + newFileName

// 		}
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"fileName": name,
// 			"url": url,
// 		})
// }