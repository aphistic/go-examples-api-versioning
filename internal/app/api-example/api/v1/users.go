package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"main/internal/pkg/logging"
)

type UsersController struct {
	logger logging.Logger
}

func NewUsersController(logger logging.Logger) *UsersController {
	return &UsersController{
		logger: logger,
	}
}

func (uc *UsersController) SetupRoutes(r chi.Router) {
	// Finally we get to the lowest level, the endpoints themselves. In this domain
	// we don't care how the users got here or the routes above us, just that
	// if they're doing a GET or POST to the root, or a GET to the /{id} route within
	// our domain that we want specific functions to be called.
	r.Get("/", uc.GetIndex)
	r.Post("/", uc.PostIndex)
	r.Get("/{id}", uc.GetUser)
}

func (uc *UsersController) GetIndex(w http.ResponseWriter, req *http.Request) {
	uc.logger.Log("Running GET / in v1.UsersController")
	w.Write([]byte("GET / v1.UsersController"))
}

func (uc *UsersController) PostIndex(w http.ResponseWriter, req *http.Request) {
	uc.logger.Log("Running POST / in v1.UsersController")
	w.Write([]byte("POST / v1.UsersController"))
}

func (uc *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	uc.logger.Log("Running GET /{id:%s} in v1.UsersController", id)
	w.Write([]byte("GET /{id:" + id + "} v1.UsersController"))
}
