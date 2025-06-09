package interfaces

import (
	"mime/multipart"
)

// CloudStorageUseCase define interface for Cloud Storage
type CloudStorageUseCase interface {
	Upload(f *multipart.FileHeader, folder string) (string, error)
	UploadSavedFile(filepath, folder string) (string, error)
	Download(path string) ([]byte, error)
	Delete(path string) error
	UploadWithSignedURL(f *multipart.FileHeader, folder string) (string, error)
	SignedURL(path string) (string, error)
}
