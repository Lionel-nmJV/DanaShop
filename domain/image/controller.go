package image

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"path/filepath"
	"strings"
)

type imageController struct {
	service  imageService
	validate *validator.Validate
}

func newController(service imageService, validate *validator.Validate) imageController {
	return imageController{
		service:  service,
		validate: validate,
	}
}

func (c imageController) uploadFile(ctx *gin.Context) {
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

	request := requestUploadImage{
		File: fileHeader,
		Type: typeImage,
		Ext:  fileExt,
	}

	image, err := NewImage().formUploadImage(request, c.validate)
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
