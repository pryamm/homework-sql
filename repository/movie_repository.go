package repository

import (
	"errors"

	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/model"
	"gorm.io/gorm"
)

func FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	db, err := config.Database()
	if err != nil {
		return nil, err
	}
	db.Find(&movies)

	return movies, nil
}

func FindOne(slug string) (model.Movie, error) {
	var movie model.Movie
	db, err := config.Database()
	if err != nil {
		return model.Movie{}, err
	}
	errNotFound := db.First(&movie, "slug = ?", slug).Error
	errors.Is(errNotFound, gorm.ErrRecordNotFound)

	return movie, errNotFound
}

func Save(movie model.Movie) (model.Movie, error) {
	db, err := config.Database()
	if err != nil {
		return model.Movie{}, err
	}
	db.Save(&movie)
	return movie, nil
}

func Delete(movie model.Movie) error {
	db, err := config.Database()
	if err != nil {
		return err
	}
	db.Delete(&movie)
	return nil
}
