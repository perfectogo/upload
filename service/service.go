package service

import (
	"github.com/perfectogo/upload/models"
	"github.com/perfectogo/upload/storage"
)

type uploadServiceInterface interface {
	UploadImg(path string) error
	GetImages() (models.PathsList, error)
}
type InterfaceServer interface {
	Upload() uploadServiceInterface
}
type Service struct {
	storage storage.InterfaceStorage
	service uploadServiceInterface
}

func NewService(storage storage.InterfaceStorage) *Service {
	return &Service{
		storage: storage,
		service: NewServiceUp(storage)}
}
func (s *Service) Upload() uploadServiceInterface {
	return s.service
}

type serviceUp struct {
	storage storage.InterfaceStorage
}

func NewServiceUp(storage storage.InterfaceStorage) *serviceUp {
	return &serviceUp{storage: storage}
}
