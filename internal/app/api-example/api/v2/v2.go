package v2

import (
	"github.com/go-chi/chi"

	v1 "main/internal/app/api-example/api/v1"
	v2users "main/internal/app/api-example/api/v2/users"
	"main/internal/pkg/group"
	"main/internal/pkg/logging"
	"main/internal/pkg/user"
)

type APIV2 struct {
	logger logging.Logger

	apiV1 *v1.APIV1

	groups *GroupsController
	users  *v2users.UsersController
}

func NewAPIV2(
	apiV1 *v1.APIV1,
	groupService *group.GroupService,
	userService *user.UserService,
	logger logging.Logger,
) *APIV2 {
	return &APIV2{
		apiV1:  apiV1,
		logger: logger,

		groups: NewGroupsController(apiV1, groupService, logger),
		users:  v2users.NewUsersController(apiV1, userService, logger),
	}
}

func (a *APIV2) SetupRoutes(r chi.Router) {
	// Continuing with the same theme as the routes before this, all this
	// domain (the V2 struct) cares about is that there's a groups controller
	// and a users controller (which happens to be in a sub-package called users),
	// and that they should be hooked up to /groups and /users respectively.
	r.Route("/groups", a.Groups().SetupRoutes)
	r.Route("/users", a.Users().SetupRoutes)
}

func (a *APIV2) Groups() *GroupsController {
	return a.groups
}

func (a *APIV2) Users() *v2users.UsersController {
	return a.users
}
