package gcs

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"

	"starter-go-gin/config"
	"starter-go-gin/utils"
)

const (
	fifteen    = 15
	twentyFour = 24
)

// GoogleCloudStorage define struct for gcs integration
type GoogleCloudStorage struct {
	cfg config.Config
}

// NewGoogleCloudStorage initiate gcs sdk
func NewGoogleCloudStorage(cfg config.Config) *GoogleCloudStorage {
	return &GoogleCloudStorage{cfg: cfg}
}

// Upload uploads file to bucket
func (g *GoogleCloudStorage) Upload(f *multipart.FileHeader, folder string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	src, err := f.Open()
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error open file")
	}

	defer func() {
		if err := src.Close(); err != nil {
			fmt.Println("error while closing gin context form file :", err)
		}
	}()

	splitFilename := strings.Split(f.Filename, ".")
	fileMime := ""

	if len(splitFilename) > 0 {
		fileMime = splitFilename[len(splitFilename)-1]
	}

	fileName := fmt.Sprintf("%s.%s", utils.SHAEncrypt(f.Filename), fileMime)

	fileStored := fmt.Sprintf("%s/%s", folder, fileName)

	sw := storageClient.Bucket(bucket).Object(fileStored).NewWriter(ctx)

	if _, err = io.Copy(sw, src); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error copy file")
	}

	if err := sw.Close(); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error close file")
	}

	u, err := url.Parse("/" + sw.Attrs().Name)
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error parse url")
	}

	return u.String(), nil
}

// UploadWithSignedURL uploads file to bucket with signed url
func (g *GoogleCloudStorage) UploadWithSignedURL(f *multipart.FileHeader, folder string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	src, err := f.Open()
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error open file")
	}

	defer func() {
		if err := src.Close(); err != nil {
			fmt.Println("error while closing gin context form file :", err)
		}
	}()

	splitFilename := strings.Split(f.Filename, ".")
	fileMime := ""

	if len(splitFilename) > 0 {
		fileMime = splitFilename[len(splitFilename)-1]
	}

	fileName := fmt.Sprintf("%s.%s", utils.SHAEncrypt(f.Filename), fileMime)

	fileStored := fmt.Sprintf("%s/%s", folder, fileName)

	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		Expires: time.Now().Add(fifteen * time.Minute),
	}

	_, err = storageClient.Bucket(bucket).SignedURL(fileStored, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}

	optsGET := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(fifteen * time.Minute),
	}

	uGET, err := storageClient.Bucket(bucket).SignedURL(fileStored, optsGET)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}

	return uGET, nil
}

// UploadSavedFile uploads file to bucket
func (g *GoogleCloudStorage) UploadSavedFile(filepath, folder string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	//nolint:gosec
	file, err := os.Open(filepath)
	if err != nil {
		return "", errors.Wrap(err, "problem opening file for gcs")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error while closing context file :", err)
		}
	}()

	splitFilename := strings.Split(file.Name(), ".")
	fileMime := ""

	if len(splitFilename) > 0 {
		fileMime = splitFilename[len(splitFilename)-1]
	}

	fileName := fmt.Sprintf("%s.%s", utils.SHAEncrypt(file.Name()), fileMime)

	fileStored := fmt.Sprintf("%s/%s", folder, fileName)

	sw := storageClient.Bucket(bucket).Object(fileStored).NewWriter(ctx)

	if _, err = io.Copy(sw, file); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error copy file")
	}

	if err := sw.Close(); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error close file")
	}

	u, err := url.Parse("/" + sw.Attrs().Name)
	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error parse url")
	}

	// Make public
	acl := storageClient.Bucket(bucket).Object(fileStored).ACL()

	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-HandleFileUploadToBucket] error while making object public")
	}

	return u.String(), nil
}

// Delete delete file from bucket
func (g *GoogleCloudStorage) Delete(path string) error {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return errors.Wrap(err, "[CloudStorageService-Delete] error get config json")
	}

	if err := storageClient.Bucket(bucket).Object(url.QueryEscape(path)).Delete(context.Background()); err != nil {
		return errors.Wrap(err, fmt.Sprintf("[CloudStorageService-Delete] unable to delete bucket %q, file %q", bucket, url.QueryEscape(path)))
	}

	return nil
}

// Download download file from bucket
func (g *GoogleCloudStorage) Download(path string) ([]byte, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "[CloudStorageService-Download] error get config json")
	}

	file, err := storageClient.Bucket(bucket).Object(path).NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[CloudStorageService-Download] error get file")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error while closing context file :", err)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "[CloudStorageService-Download] error read file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			// Log or handle the error appropriately
			fmt.Println("Error closing file:", err)
		}
	}()

	return data, nil
}

// SignedURL get signed url
func (g *GoogleCloudStorage) SignedURL(path string) (string, error) {
	bucket := g.cfg.Google.StorageBucketName

	var err error

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {
		return "", errors.Wrap(err, "[CloudStorageService-Upload] error get config json")
	}

	optsGET := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(twentyFour * time.Hour),
	}

	// remove slash
	path = strings.TrimPrefix(path, "/")

	uGET, err := storageClient.Bucket(bucket).SignedURL(path, optsGET)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}

	return uGET, nil
}
