package middlewares

import (
	"fmt"
	"net/http"
	"path/filepath"
	"personal-growth/utils"

	"github.com/gofiber/fiber/v2"
)

type UploadFileOptions struct {
	AllowedTypes map[string]bool
	FileSize     int64
	BasePath     string
}

func Uploadfile(options UploadFileOptions) fiber.Handler {
	if options.FileSize == 0 {
		options.FileSize = 5
	}

	return func(c *fiber.Ctx) error {
		// Lấy file
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		// limit file size to 5MB
		if file.Size > options.FileSize*1024*1024 {
			return c.Status(400).SendString(fmt.Sprintf("File size is too big, Maximum is %dMB", options.FileSize))
		}

		//check file typetype
		f, err := file.Open()
		if err != nil {
			return c.Status(500).SendString("Can not open file")
		}
		defer f.Close()
		buf := make([]byte, 512)
		_, err = f.Read(buf)
		if err != nil {
			return c.Status(500).SendString("Can not read file")
		}

		mimeType := http.DetectContentType(buf)
		if !options.AllowedTypes[mimeType] {
			return c.Status(400).SendString("Invalid file type")
		}

		//edit file name
		ext := filepath.Ext(file.Filename)
		newFileName := utils.GenerateFilename(ext)

		// Lưu file vào thư mục uploads/
		utils.EnsureDir("./uploads/avatar")
		filepath := fmt.Sprintf("./uploads/%s/%s", options.BasePath, newFileName)
		err = c.SaveFile(file, filepath)
		if err != nil {
			return err
		}

		// Gắn user_id vào context để dùng trong handler
		c.Locals("file", filepath[1:])

		return c.Next()
	}
}
