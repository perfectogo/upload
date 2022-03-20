package repository

import "github.com/perfectogo/upload/models"

type UploadFileInterface interface {
	UploadImg(path string) error
	GetImages() (models.PathsList, error)
}
