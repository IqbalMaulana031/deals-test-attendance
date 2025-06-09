package service

import (
	"context"
	"mime/multipart"

	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
)

// UtilsCreator is a creator for utils
type UtilsCreator struct {
	cfg          config.Config
	cloudStorage interfaces.CloudStorageUseCase
}

// UtilsCreatorUseCase is a use case for utils
type UtilsCreatorUseCase interface {
	// UploadFile is a function that uploads a file
	UploadFile(ctx context.Context, file []*multipart.FileHeader, folder string) ([]string, error)
}

// NewUtilsCreator is a constructor for UtilsCreator
func NewUtilsCreator(
	cfg config.Config,
	cloudStorage interfaces.CloudStorageUseCase,
) *UtilsCreator {
	return &UtilsCreator{
		cfg:          cfg,
		cloudStorage: cloudStorage,
	}
}
