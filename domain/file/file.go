package file

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"starfish/config"
	"strconv"
)

type File struct {
	File *multipart.FileHeader `json:"file"validation:"required"`
	Type string                `json:"type"validation:"required"`
}

func NewFile() File {
	return File{}
}

func (f File) formUploadImage(request requestUploadFile, validate *validator.Validate) (File, error) {
	err := validate.Struct(request)
	if err != nil {
		return f, errors.New("invalid request")
	}

	if request.Type != "campaign" && request.Ext == "mp4" || request.Type == "campaign" && request.Ext != "mp4" {
		return f, errors.New("invalid request")
	}

	f.File = request.File
	f.Type = request.Type

	return f, nil
}

func (f File) validateSize(request requestUploadFile, config config.CloudinaryConfig) error {
	imageSize, err := strconv.Atoi(config.ImageSize)
	if err != nil {
		return errors.New("invalid request")
	}
	videoSize, err := strconv.Atoi(config.VideoSize)
	if err != nil {
		return errors.New("invalid request")
	}

	if request.Ext == "mp4" && request.File.Size > (int64(videoSize)*1024*1024) {
		return errors.New("invalid request")
	} else if request.Ext != "mp4" && request.File.Size > (int64(imageSize)*1024*1024) {
		return errors.New("invalid request")
	}
	return nil
}
