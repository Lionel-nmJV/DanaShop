package file

import "mime/multipart"

type requestUploadFile struct {
	File *multipart.FileHeader `json:"file"validate:"required"`
	Type string                `json:"type"validate:"required,oneof=product merchant campaign"`
	Ext  string                `json:"ext"validate:"required,oneof=png jpg jpeg webp mp4"`
}
