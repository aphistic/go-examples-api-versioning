package api

import (
	"github.com/go-chi/chi"

	v1 "main/internal/app/api-example/api/v1"
	v2 "main/internal/app/api-example/api/v2"
	"main/internal/pkg/group"
	"main/internal/pkg/logging"
	"main/internal/pkg/user"
)

type Option func(api *API)

func WithLogger(logger logging.Logger) Option {
	return func(api *API) {
		api.logger = logger
	}
}

type API struct {
	logger logging.Logger

	groupService *group.GroupService
	userService  *user.UserService
}

func NewAPI(
	groupService *group.GroupService,
	userService *user.UserService,
	opts ...Option,
) *API {
	a := &API{
		logger: logging.NewNilLogger(),

		userService:  userService,
		groupService: groupService,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *API) SetupRoutes(r chi.Router) {
	// By doing the setup for the versions here instead of a higher level
	// we have full control over them without needing the higher domains (or the
	// world outside of the API) to know anything about the internals.

	// The domain that the API knows about in this package is that there's something
	// called a v1 API and something called a v2 API, but doesn't know (or care!) what's
	// inside.

	// All this domain cares about is that the v1 API should be on /v1 and the v2 API
	// should be on /v2.
	apiV1 := v1.NewAPIV1(
		a.groupService,
		a.userService,
		a.logger,
	)
	r.Route("/v1", apiV1.SetupRoutes)

	apiV2 := v2.NewAPIV2(
		apiV1,
		a.groupService,
		a.userService,
		a.logger,
	)
	r.Route("/v2", apiV2.SetupRoutes)
}
