package v1

import (
	"github.com/go-chi/chi"

	"main/internal/pkg/logging"
)

type APIV1 struct {
	logger logging.Logger

	groups *GroupsController
	users  *UsersController
}

func NewAPIV1(logger logging.Logger) *APIV1 {
	return &APIV1{
		logger: logger,

		groups: NewGroupsController(logger),
		users:  NewUsersController(logger),
	}
}

func (a *APIV1) SetupRoutes(r chi.Router) {
	// Continuing with the same theme as the routes before this, all this
	// domain (the V1 struct) cares about is that there's a groups controller
	// and a users controller and that they should be hooked up to /groups and /users
	// respectively.
	r.Route("/groups", a.Groups().SetupRoutes)
	r.Route("/users", a.Users().SetupRoutes)
}

func (a *APIV1) Groups() *GroupsController {
	return a.groups
}

func (a *APIV1) Users() *UsersController {
	return a.users
}
