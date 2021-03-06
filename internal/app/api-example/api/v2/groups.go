package v2

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "main/internal/app/api-example/api/v1"
	v1models "main/internal/app/api-example/api/v1/models"
	"main/internal/app/api-example/encoding"
	"main/internal/pkg/errors"
	"main/internal/pkg/group"
	"main/internal/pkg/logging"
)

type GroupsController struct {
	logger logging.Logger

	apiV1 *v1.APIV1

	groupService *group.GroupService
}

func NewGroupsController(
	apiV1 *v1.APIV1,
	groupService *group.GroupService,
	logger logging.Logger,
) *GroupsController {
	return &GroupsController{
		logger:       logger,
		apiV1:        apiV1,
		groupService: groupService,
	}
}

func (gc *GroupsController) SetupRoutes(r chi.Router) {
	// Finally we get to the lowest level, the endpoints themselves. In this domain
	// we don't care how the users got here or the routes above us, just that
	// if they're doing a GET or POST to the root, or a GET to the /{id} route within
	// our domain that we want specific functions to be called.

	// This is very similar to the v1 controller but we can take advantage of the fact that
	// our v2 api hasn't changed GET / and POST / in a breaking way so we can just call those
	// methods on v1. GET /{id} has changed, though, so we need to implement that.
	r.Get("/", gc.apiV1.Groups().GetIndex)
	r.Post("/", gc.apiV1.Groups().PostIndex)
	r.Get("/{id}", gc.GetGroup)
}

func (gc *GroupsController) GetGroup(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	gc.logger.Log("Running GET /{id:%s} in v2.GroupsController", id)

	reqGroup, err := gc.groupService.GetGroupByID(id)
	if err != nil && err == errors.ErrNotFound {
		err = encoding.WriteJSONResult(
			w, http.StatusNotFound,
			&v1models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	} else if err != nil {
		gc.logger.Log("Could not get group: %s", err)
		err = encoding.WriteJSONResult(
			w, http.StatusInternalServerError,
			&v1models.ErrorResponse{Error: err.Error()},
		)
		if err != nil {
			gc.logger.Log("Could not write error message: %s", err)
		}
		return
	}

	err = encoding.WriteJSONResult(
		w, http.StatusOK,
		&GroupDetailResponse{
			ID:   reqGroup.ID,
			Name: reqGroup.Name,
		},
	)
	if err != nil {
		gc.logger.Log("Could not write response: %s\n", err)
		return
	}

	w.Write([]byte("\nGET /{id:" + id + "} v2.GroupsController\n"))
}
