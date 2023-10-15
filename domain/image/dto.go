package image

import "mime/multipart"

type requestUploadImage struct {
	File *multipart.FileHeader `json:"file"validate:"required"`
	Type string                `json:"type"validate:"required"`
	Ext  string                `json:"ext"validate:"required,oneof=png jpg jpeg webp"`
}
