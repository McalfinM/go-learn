package storage

import "mime/multipart"

type Storage interface {
	Upload(file *multipart.FileHeader, folder string) (string, error)
	Delete(url string) error
}
