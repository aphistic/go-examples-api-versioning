package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"main/internal/pkg/logging"
	"main/internal/pkg/user"
)

type UsersController struct {
	logger logging.Logger

	userService *user.UserService
}

func NewUsersController(userService *user.UserService, logger logging.Logger) *UsersController {
	return &UsersController{
		logger:      logger,
		userService: userService,
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

	users, err := uc.userService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET / v1.UsersController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Display users somehow
	_ = users

	w.Write([]byte("GET / v1.UsersController\n"))
}

func (uc *UsersController) PostIndex(w http.ResponseWriter, req *http.Request) {
	uc.logger.Log("Running POST / in v1.UsersController")

	err := uc.userService.CreateUser(&user.User{
		ID:    "generate somehow",
		Email: "from post",
		Name:  "from post",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("POST / v1.UsersController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("POST / v1.UsersController\n"))
}

func (uc *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	uc.logger.Log("Running GET /{id:%s} in v1.UsersController", id)

	reqUser, err := uc.userService.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET /{id:" + id + "} v1.UsersController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Do something with the user
	_ = reqUser

	w.Write([]byte("GET /{id:" + id + "} v1.UsersController\n"))
}
