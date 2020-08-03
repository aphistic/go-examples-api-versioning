package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"main/internal/app/api-example/api"
	"main/internal/pkg/logging"
)

func main() {
	// Start out by creating a stdout logger for ourselves to use
	logger := logging.NewStdoutLogger()

	// Create a chi router that we'll use to handle all the incoming
	// traffic and route it to the right places
	r := chi.NewRouter()

	// Create a root-level API object with our logger
	httpAPI := api.NewAPI(
		api.WithLogger(logger),
	)
	// Tell chi to route our API from the /api route. Since the rest of the
	// API setup is handled in the nested structs we don't need to setup anything
	// else here. A benefit of organizing the code this way is we can move the route
	// anywhere without needing to change anything lower in the route.
	r.Route("/api", httpAPI.SetupRoutes)
	// We could even easily add another route if we happened to want to rename the root
	// route just by adding another route call:
	// r.Route("/notapi", httpAPI.SetupRoutes)

	// From here on it's a standard Go http.Server. Set the address to listen on
	// and give the server the chi router to handle the incoming requests.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	logger.Log("Starting http server")
	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		// Ignore server closed because it's expected
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "http server error: %s\n", err)
		os.Exit(1)
	}
}
