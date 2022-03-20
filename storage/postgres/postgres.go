package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/perfectogo/upload/models"
)

type uploadRepo struct {
	db *sqlx.DB
}

func NewUploadRepo(db *sqlx.DB) *uploadRepo {
	return &uploadRepo{db: db}
}
func (r *uploadRepo) UploadImg(path string) error {
	if err := r.db.QueryRow("Insert into images (imgURL) values ($1)", path); err != nil {
		return err.Err()
	}
	return nil
}
func (r *uploadRepo) GetImages() (models.PathsList, error) {
	rows, err := r.db.Queryx("select imgURL from images")
	if err != nil {
		return models.PathsList{}, err
	}
	defer rows.Close()
	var (
		images []*models.Path
		count  int64
	)
	for rows.Next() {
		var image models.Path
		rows.Scan(&image.Path)
		images = append(images, &image)
	}
	if err = r.db.QueryRow(`select count(*) from images`).Scan(&count); err != nil {
		return models.PathsList{}, err
	}
	log.Println(&images)
	return models.PathsList{
		Paths: images,
		Count: count}, nil
}
