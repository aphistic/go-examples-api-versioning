package site

import (
	"github.com/go-chi/chi"
)

// Site represents non-API portions of the site.
type Site struct{}

func NewSite() *Site {
	return &Site{}
}

func (s *Site) SetupRoutes(r chi.Router) {
	r.Route("/", newHomeController().SetupRoutes)
}
