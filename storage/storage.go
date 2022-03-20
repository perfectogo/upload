package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/perfectogo/upload/storage/postgres"
	"github.com/perfectogo/upload/storage/repository"
)

type InterfaceStorage interface {
	Upload() repository.UploadFileInterface
}

type storagePg struct {
	db         *sqlx.DB
	uploadRepo repository.UploadFileInterface
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		uploadRepo: postgres.NewUploadRepo(db)}
}
func (s *storagePg) Upload() repository.UploadFileInterface {
	return s.uploadRepo
}
