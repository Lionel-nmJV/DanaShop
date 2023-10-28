package file

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"path/filepath"
	"starfish/config"
	"strings"
)

type fileController struct {
	service  fileService
	validate *validator.Validate
	config   config.CloudinaryConfig
}

func newController(service fileService, validate *validator.Validate, config config.CloudinaryConfig) fileController {
	return fileController{
		service:  service,
		validate: validate,
		config:   config,
	}
}

func (c fileController) uploadFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		writeError(ctx, errors.New("invalid request"), 40001, http.StatusBadRequest)
		return
	}

	typeImage := ctx.PostForm("type")
	if err != nil {
		writeError(ctx, errors.New("invalid request"), 40001, http.StatusBadRequest)
		return
	}

	fileExt := filepath.Ext(fileHeader.Filename)
	fileExt = strings.TrimPrefix(fileExt, ".")

	request := requestUploadFile{
		File: fileHeader,
		Type: typeImage,
		Ext:  fileExt,
	}

	image, err := NewFile().formUploadImage(request, c.validate)
	if err != nil {
		writeError(ctx, err, 40002, http.StatusBadRequest)
		return
	}

	err = NewFile().validateSize(request, c.config)
	if err != nil {
		writeError(ctx, err, 40002, http.StatusBadRequest)
		return
	}

	secureURL, err := c.service.uploadImage(ctx, image)
	if err != nil {
		writeError(ctx, err, 40001, http.StatusBadRequest)
		return
	}

	writeSuccess(ctx, gin.H{"url": secureURL}, http.StatusOK)
}
