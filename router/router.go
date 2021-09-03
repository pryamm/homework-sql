package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/model"
	"github.com/rysmaadit/go-template/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()
	db, _ := config.Database()
	db.AutoMigrate(
		&model.Movie{},
	)
	setAuthRouter(r, dependencies.AuthService)
	setMovieRouter(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func setMovieRouter(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/movie").Handler(handler.GetAllMovie())
	router.Methods(http.MethodGet).Path("/movie/{slug}").Handler(handler.GetOneMovie())
	router.Methods(http.MethodPost).Path("/movie").Handler(handler.CreateMovie())
	router.Methods(http.MethodPut).Path("/movie/{slug}").Handler(handler.UpdateMovie())
	router.Methods(http.MethodDelete).Path("/movie/{slug}").Handler(handler.DeleteMovie())
}
