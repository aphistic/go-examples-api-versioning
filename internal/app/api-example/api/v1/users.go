package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"main/internal/app/api-example/api/v1/models"
	"main/internal/app/api-example/encoding"
	"main/internal/pkg/errors"
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
		uc.logger.Log("Could not get user list: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			uc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	// Translate to our view model
	result := make(models.ListUsersResponse, 0, len(users))
	for _, resultUser := range users {
		result = append(result, &models.UserResponse{
			ID:    resultUser.ID,
			Email: resultUser.Email,
		})
	}

	err = encoding.WriteJSONResult(w, http.StatusOK, result)
	if err != nil {
		uc.logger.Log("Could not write response: %s\n", err)
		return
	}
	w.Write([]byte("\nGET / v1.UsersController\n"))
}

func (uc *UsersController) PostIndex(w http.ResponseWriter, req *http.Request) {
	uc.logger.Log("Running POST / in v1.UsersController")

	generatedID := "generate somehow"
	err := uc.userService.CreateUser(&user.User{
		ID:    generatedID,
		Email: "from post",
		Name:  "from post",
	})
	if err != nil {
		uc.logger.Log("Could not create user: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			&models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			uc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	err = encoding.WriteJSONResult(
		w, http.StatusOK,
		&models.CreateUserResponse{ID: generatedID},
	)
	if err != nil {
		uc.logger.Log("Could not write response: %s\n", err)
		return
	}

	w.Write([]byte("\nPOST / v1.UsersController\n"))
}

func (uc *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	uc.logger.Log("Running GET /{id:%s} in v1.UsersController", id)

	reqUser, err := uc.userService.GetUserByID(id)
	if err != nil && err == errors.ErrNotFound {
		err = encoding.WriteJSONResult(
			w, http.StatusNotFound,
			&models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			uc.logger.Log("Could not write error message: %s", err)
		}
		return
	} else if err != nil {
		uc.logger.Log("Could not get user: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			&models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			uc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	err = encoding.WriteJSONResult(
		w, http.StatusOK,
		&models.UserDetailResponse{
			ID:    reqUser.ID,
			Email: reqUser.Email,
			Name:  reqUser.Name,
		},
	)
	if err != nil {
		uc.logger.Log("Could not write response: %s\n", err)
		return
	}

	w.Write([]byte("\nGET /{id:" + id + "} v1.UsersController\n"))
}
