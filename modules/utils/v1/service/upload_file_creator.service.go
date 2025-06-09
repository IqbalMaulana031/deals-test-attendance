package service

import (
	"context"
	"mime/multipart"

	"starter-go-gin/common/errors"
	"starter-go-gin/common/logger"
)

// UploadFile is a function that uploads a file
func (ufc *UtilsCreator) UploadFile(ctx context.Context, file []*multipart.FileHeader, folder string) ([]string, error) {
	if folder == "" {
		folder = "pertamina-files"
	}

	paths := []string{}
	for _, f := range file {
		if f.Size > 10<<20 {
			logger.ErrorFromStr(ctx, "File size is too big")
			return nil, errors.ErrFileMaxSize.Error()
		}
		filePath, err := ufc.cloudStorage.Upload(f, folder)
		if err != nil {
			logger.Error(ctx, err)
			return nil, errors.ErrInternalServerError.Error()
		}
		paths = append(paths, filePath)
	}

	return paths, nil
}
