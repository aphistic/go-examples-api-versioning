package v1

import (
	"net/http"

	"github.com/go-chi/chi"

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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET / v1.GroupsController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Do what we need to with the groups!
	_ = groups

	w.Write([]byte("GET / v1.GroupsController\n"))
}

func (gc *GroupsController) PostIndex(w http.ResponseWriter, req *http.Request) {
	gc.logger.Log("Running POST / in v1.GroupsController")

	err := gc.groupService.CreateGroup(&group.Group{
		ID:   "generate",
		Name: "from post data",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("POST / v1.GroupsController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("POST / v1.GroupsController\n"))
}

func (gc *GroupsController) GetGroup(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	gc.logger.Log("Running GET /{id:%s} in v1.GroupsController", id)

	reqGroup, err := gc.groupService.GetGroupByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("GET /{id:" + id + "} v1.GroupsController\n"))
		w.Write([]byte(err.Error()))
		return
	}

	// Do what we need to for returning the group
	_ = reqGroup

	w.Write([]byte("GET /{id:" + id + "} v1.GroupsController\n"))
}
