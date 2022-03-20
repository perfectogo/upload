package service

import (
	"github.com/perfectogo/upload/models"
)

func (s *serviceUp) UploadImg(path string) error {
	if err := s.storage.Upload().UploadImg(path); err != nil {
		return err
	}
	return nil
}
func (s *serviceUp) GetImages() (models.PathsList, error) {
	images, err := s.storage.Upload().GetImages()
	if err != nil {
		return models.PathsList{}, err
	}
	return images, nil
}
