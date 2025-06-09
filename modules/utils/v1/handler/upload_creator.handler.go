package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"starter-go-gin/common/errors"
	"starter-go-gin/resource"
	"starter-go-gin/response"
)

// UploadFile is a handler for upload file
func (u *UtilsCreatorHandler) UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()

	files := form.File["file"]

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	filePaths, err := u.creator.UploadFile(c, files, c.PostForm("folder"))

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	baseURL := u.cfg.Google.StorageEndpoint

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.UploadFile{
		Path:    filePaths,
		BaseURL: baseURL,
	}))
}
