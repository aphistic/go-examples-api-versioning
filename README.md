# API Versioning

This is an example of how I like to lay out versioned (or even unversioned) APIs. This example
is centered around a RESTful, versioned HTTP API but the ideas can be used elsewhere.

The general idea is that each "domain" of the API doesn't need to be aware of any other domain,
just itself. The top level API, for example, doesn't need to know about the structure of the code
OR the routes for another package. It follows a pattern similar to MVC... but just the C part. The
MV part is up to you. :)

For the code from a v2 api that is the same as the v1 API (it hasn't changed and functions exactly
the same) I haven't really come up with a pattern that I'm super happy with. The best idea I've had
so far is to pass an instance of the previous version to the next version and having the next version
call the previous version. You'll see this in `internal/app/api-example/api/v2`.

Here's a layout of the API itself:

```
+ /api
  + /v1
    + /users
      - GET
      - POST
      - GET /{id}
    + /groups
      - GET
      - POST
      - GET /{id}
  + /v2
    + /users
      - GET
      - POST
      - GET /{id}
    + /groups
      - GET
      - POST
      - GET /{id}
```

The repo is laid out similar to [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
as I find it's the most extensive, well-documented and accepted example of laying out a Go
binary project.

The best place to start reading this code is `cmd/api-example/main.go` because that's where the
initial API and HTTP server are started.