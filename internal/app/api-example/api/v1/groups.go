package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"main/internal/app/api-example/api/v1/models"
	"main/internal/app/api-example/encoding"
	"main/internal/pkg/errors"
	"main/internal/pkg/group"
	"main/internal/pkg/logging"
)

type GroupsController struct {
	logger logging.Logger

	groupService *group.GroupService
}

func NewGroupsController(groupService *group.GroupService, logger logging.Logger) *GroupsController {
	return &GroupsController{
		logger:       logger,
		groupService: groupService,
	}
}

func (gc *GroupsController) SetupRoutes(r chi.Router) {
	// Finally we get to the lowest level, the endpoints themselves. In this domain
	// we don't care how the users got here or the routes above us, just that
	// if they're doing a GET or POST to the root, or a GET to the /{id} route within
	// our domain that we want specific functions to be called.
	r.Get("/", gc.GetIndex)
	r.Post("/", gc.PostIndex)
	r.Get("/{id}", gc.GetGroup)
}

func (gc *GroupsController) GetIndex(w http.ResponseWriter, req *http.Request) {
	gc.logger.Log("Running GET / in v1.GroupsController")

	groups, err := gc.groupService.ListGroups()
	if err != nil {
		gc.logger.Log("Could not get group list: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	// Translate to our view model
	result := make(models.ListGroupsResponse, 0, len(groups))
	for _, resultGroup := range groups {
		result = append(result, &models.GroupResponse{
			ID:   resultGroup.ID,
			Name: resultGroup.Name,
		})
	}

	err = encoding.WriteJSONResult(w, http.StatusOK, result)
	if err != nil {
		gc.logger.Log("Could not write response: %s\n", err)
		return
	}
	w.Write([]byte("\nGET / v1.GroupsController\n"))
}

func (gc *GroupsController) PostIndex(w http.ResponseWriter, req *http.Request) {
	gc.logger.Log("Running POST / in v1.GroupsController")

	generatedID := "generate somehow"
	err := gc.groupService.CreateGroup(&group.Group{
		ID:   generatedID,
		Name: "from post data",
	})
	if err != nil {
		gc.logger.Log("Could not create group: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	err = encoding.WriteJSONResult(
		w, http.StatusOK,
		&models.CreateGroupResponse{ID: generatedID},
	)
	if err != nil {
		gc.logger.Log("Could not write response: %s\n", err)
		return
	}

	w.Write([]byte("\nPOST / v1.GroupsController\n"))
}

func (gc *GroupsController) GetGroup(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	gc.logger.Log("Running GET /{id:%s} in v1.GroupsController", id)

	reqGroup, err := gc.groupService.GetGroupByID(id)
	if err != nil && err == errors.ErrNotFound {
		err = encoding.WriteJSONResult(
			w, http.StatusNotFound,
			&models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	} else if err != nil {
		gc.logger.Log("Could not get group: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			&models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	err = encoding.WriteJSONResult(
		w, http.StatusOK,
		&models.GroupDetailResponse{
			ID:   reqGroup.ID,
			Name: reqGroup.Name,
		},
	)
	if err != nil {
		gc.logger.Log("Could not write response: %s\n", err)
		return
	}

	w.Write([]byte("\nGET /{id:" + id + "} v1.GroupsController\n"))
}
