package storage

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryStorage struct {
	client *cloudinary.Cloudinary
}

func NewCloudinaryStorage() (*CloudinaryStorage, error) {
	cld, err := cloudinary.NewFromParams(
		"YOUR_CLOUD_NAME",
		"YOUR_API_KEY",
		"YOUR_API_SECRET",
	)
	if err != nil {
		return nil, err
	}

	return &CloudinaryStorage{client: cld}, nil
}

func (c *CloudinaryStorage) Upload(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	res, err := c.client.Upload.Upload(context.Background(), src, uploader.UploadParams{
		Folder: folder,
	})

	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}

func (c *CloudinaryStorage) Delete(url string) error {
	// optional (butuh public_id extraction)
	return nil
}
