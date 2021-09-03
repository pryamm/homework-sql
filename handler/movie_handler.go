package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/model"
	"github.com/rysmaadit/go-template/service"
)

func GetAllMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := service.FindAllMovie()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, movies, nil)
	}
}

func GetOneMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		movie, err := service.FindOneMovie(slug)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
			return
		}
		responder.NewHttpResponse(r, w, http.StatusOK, movie, nil)
	}

}

func CreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := new(model.Movie)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		defer r.Body.Close()

		movie, errorInsert := service.CreateMovie(*payload)
		if errorInsert != nil {
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, errorInsert)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusCreated, movie, nil)
	}
}

func UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		payload := new(model.Movie)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)

		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		update, errorUpdate := service.UpdateMovie(slug, *payload)
		if errorUpdate != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, errorUpdate)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, update, nil)
	}
}

func DeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		_, errorDeleteMovie := service.DeleteMovie(slug)
		if errorDeleteMovie != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, errorDeleteMovie)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, nil, nil)
	}
}
