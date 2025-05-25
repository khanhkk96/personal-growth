package middlewares

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"personal-growth/utils"

	"github.com/gofiber/fiber/v2"
)

type UploadFileOptions struct {
	AllowedTypes     map[string]bool
	FileSize         int64
	Destination      string
	MaximumFileCount int
}

func validateFileType(file *multipart.FileHeader, allowedTypes map[string]bool) bool {
	//check file typetype
	f, err := file.Open()
	if err != nil {
		fmt.Println("Can not open file")
		return false
	}
	defer f.Close()
	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Can not read file")
		return false
	}

	mimeType := http.DetectContentType(buf)
	if !allowedTypes[mimeType] {
		fmt.Println("Invalid file type")
		return false
	}

	return true
}

func UploadFileHandlder(options UploadFileOptions) fiber.Handler {
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

		//validate file type
		if ok := validateFileType(file, options.AllowedTypes); !ok {
			return c.Status(400).SendString("Invalid file type")
		}

		//edit file name
		ext := filepath.Ext(file.Filename)
		newFileName := utils.GenerateFilename(ext)

		// Lưu file vào thư mục uploads/
		utils.EnsureDir(fmt.Sprintf("./uploads/%s", options.Destination))
		filepath := fmt.Sprintf("./uploads/%s/%s", options.Destination, newFileName)
		err = c.SaveFile(file, filepath)
		if err != nil {
			return err
		}

		// Gắn user_id vào context để dùng trong handler
		c.Locals("file", filepath[1:])

		return c.Next()
	}
}

func UploadMultiFilesHandlder(options UploadFileOptions) fiber.Handler {
	if options.FileSize == 0 {
		options.FileSize = 5
	}

	if options.MaximumFileCount == 0 {
		options.MaximumFileCount = 5
	}

	return func(c *fiber.Ctx) error {
		// Lấy file
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := form.File["files"]

		if len(files) > options.MaximumFileCount {
			return c.Status(400).SendString("Too many files")
		}

		fileList := []string{}

		for _, file := range files {
			// limit file size to 5MB
			if file.Size > options.FileSize*1024*1024 {
				return c.Status(400).SendString(fmt.Sprintf("File size is too big, Maximum is %dMB", options.FileSize))
			}

			//validate file type
			if ok := validateFileType(file, options.AllowedTypes); !ok {
				return c.Status(400).SendString("Invalid file type")
			}

			//edit file name
			ext := filepath.Ext(file.Filename)
			newFileName := utils.GenerateFilename(ext)

			// Lưu file vào thư mục uploads/
			utils.EnsureDir(fmt.Sprintf("./uploads/%s", options.Destination))
			filepath := fmt.Sprintf("./uploads/%s/%s", options.Destination, newFileName)
			err = c.SaveFile(file, filepath)
			if err != nil {
				return err
			}

			fileList = append(fileList, filepath[1:])
		}

		// Gắn user_id vào context để dùng trong handler
		c.Locals("files", fileList)

		return c.Next()
	}
}
