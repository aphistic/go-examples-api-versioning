package site

import (
	"net/http"

	"github.com/go-chi/chi"
)

// homeController handles any top-level home-type user requests
type homeController struct{}

func newHomeController() *homeController {
	return &homeController{}
}

func (hc *homeController) SetupRoutes(r chi.Router) {
	r.Get("/", hc.getIndex)
}

// getIndex is the index page for our entire site
func (hc *homeController) getIndex(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`
<h1>Go API Versioning Example</h1>
<br />
Take a look at <a href="https://repl.it/@aphistic/APIVersioning">the code</a> to see 
how this is laid out.<br />
<br />

<h2>Example Requests</h2>
<h3>V1 Users</h3>
<a href="/api/v1/users">/api/v1/users</a><br />
<a href="/api/v1/users/1234">/api/v1/users/1234</a><br />
<h3>V1 Groups</h3>
<a href="/api/v1/groups">/api/v1/groups</a><br />
<a href="/api/v1/groups/1234">/api/v1/groups/1234</a><br />
<h3>V2 Users</h3>
<a href="/api/v2/users">/api/v2/users</a><br />
<a href="/api/v2/users/1234">/api/v2/users/1234</a><br />
<h3>V2 Groups</h3>
<a href="/api/v2/groups">/api/v2/groups</a><br />
<a href="/api/v2/groups/1234">/api/v2/groups/1234</a><br />

<h2>Example cURL Commands</h2>
<h3>V1 Users</h3>
<code>$ curl http://localhost:8080/api/v1/users</code><br />
<code>$ curl -X POST http://localhost:8080/api/v1/users</code><br />
<code>$ curl http://localhost:8080/api/v1/users/1234</code><br />
<h3>V1 Groups</h3>
<code>$ curl http://localhost:8080/api/v1/groups</code><br />
<code>$ curl -X POST http://localhost:8080/api/v1/groups</code><br />
<code>$ curl http://localhost:8080/api/v1/groups/1234</code><br />
<h3>V2 Users</h3>
<code>$ curl http://localhost:8080/api/v2/users</code><br />
<code>$ curl -X POST http://localhost:8080/api/v2/users</code><br />
<code>$ curl http://localhost:8080/api/v2/users/1234</code><br />
<h3>V2 Groups</h3>
<code>$ curl http://localhost:8080/api/v2/groups</code><br />
<code>$ curl -X POST http://localhost:8080/api/v2/groups</code><br />
<code>$ curl http://localhost:8080/api/v2/groups/1234</code><br />
	`))
}
