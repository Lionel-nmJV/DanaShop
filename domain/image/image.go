package image

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
)

type Image struct {
	File *multipart.FileHeader `json:"file"validation:"required"`
	Type string                `json:"type"validation:"required"`
}

func NewImage() Image {
	return Image{}
}

func (i Image) formUploadImage(request requestUploadImage, validate *validator.Validate) (Image, error) {
	err := validate.Struct(request)
	if err != nil {
		return i, errors.New("invalid file type")
	}
	i.File = request.File
	i.Type = request.Type

	return i, nil
}
