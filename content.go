package core

import (
	"mime/multipart"
)

type MultipartFile struct {
	File   multipart.File
	Header *multipart.FileHeader
}
