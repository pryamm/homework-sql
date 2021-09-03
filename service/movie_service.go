package service

import (
	"github.com/rysmaadit/go-template/model"
	"github.com/rysmaadit/go-template/repository"
)

func FindAllMovie() ([]model.Movie, error) {
	result, err := repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindOneMovie(slug string) (model.Movie, error) {
	result, errSlug := repository.FindOne(slug)
	if errSlug != nil {
		return model.Movie{}, errSlug
	}
	return result, nil
}

func CreateMovie(movie model.Movie) (model.Movie, error) {
	inserted, err := repository.Save(movie)
	if err != nil {
		return model.Movie{}, nil
	}
	return inserted, nil
}

func UpdateMovie(slug string, payload model.Movie) (model.Movie, error) {
	movie, errSlug := repository.FindOne(slug)
	if errSlug != nil {
		return model.Movie{}, errSlug
	}

	movie.Title = payload.Title
	movie.Slug = payload.Slug
	movie.Description = payload.Description
	movie.Duration = payload.Duration
	movie.Image = payload.Image

	update, errUpdate := repository.Save(movie)
	if errUpdate != nil {
		return model.Movie{}, errUpdate
	}

	return update, nil
}

func DeleteMovie(slug string) (interface{}, error) {
	movie, errSlug := repository.FindOne(slug)
	if errSlug != nil {
		return nil, errSlug
	}

	errDelete := repository.Delete(movie)
	if errDelete != nil {
		return nil, errDelete
	}

	return nil, nil
}
