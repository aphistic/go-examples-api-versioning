// Package users exists to show that the controllers can be broken
// out to their own packages if needed.
package users

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "main/internal/app/api-example/api/v1"
	"main/internal/pkg/logging"
	"main/internal/pkg/user"
)

type UsersController struct {
	logger logging.Logger

	apiV1 *v1.APIV1

	userService *user.UserService
}

func NewUsersController(
	apiV1 *v1.APIV1,
	userService *user.UserService,
	logger logging.Logger,
) *UsersController {
	return &UsersController{
		logger:      logger,
		apiV1:       apiV1,
		userService: userService,
	}
}

func (uc *UsersController) SetupRoutes(r chi.Router) {
	// Finally we get to the lowest level, the endpoints themselves. In this domain
	// we don't care how the users got here or the routes above us, just that
	// if they're doing a GET or POST to the root, or a GET to the /{id} route within
	// our domain that we want specific functions to be called.

	// This is very similar to the v1 controller but we can take advantage of the fact that
	// our v2 api hasn't changed POST / in a breaking way so we can just call those
	// methods on v1. GET / and GET /{id} have changed, though, so we need to implement them.
	r.Get("/", uc.GetIndex)
	r.Post("/", uc.apiV1.Users().PostIndex)
	r.Get("/{id}", uc.GetUser)
}

func (uc *UsersController) GetIndex(w http.ResponseWriter, req *http.Request) {
	uc.logger.Log("Running GET / in v2users.UsersController")

	users, err := uc.userService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET / v2users.UsersController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Do what we need to with the users
	_ = users

	w.Write([]byte("GET / v2users.UsersController\n"))
}

func (uc *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	uc.logger.Log("Running GET /{id:%s} in v2users.UsersController", id)

	reqUser, err := uc.userService.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET /{id:" + id + "} v2users.UsersController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Do what we need to with the user
	_ = reqUser

	w.Write([]byte("GET /{id:" + id + "} v2users.UsersController\n"))
}
