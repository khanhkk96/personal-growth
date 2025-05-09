package constants

var ImageFileTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/jpg":  true,
	"image/gif":  true,
	"image/webp": true,
}

var DocFileTypes = map[string]bool{
	"application/pdf":    true,
	"text/plain":         true,
	"application/msword": true,
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/vnd.ms-excel": true,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true,
}

var ExcelFileTypes = map[string]bool{
	"application/vnd.ms-excel": true,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true,
}
