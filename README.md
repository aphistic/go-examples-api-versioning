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

The models of the data returned from the API are shown in a couple different example variations
depending on how you'd like to handle it. The `v1` models are in a subpackage of the `v1` package
and the `v2` models are in files next to the controllers themselves. I've found both these patterns
work as long as you're consistent. Having the models in a separate package can help organize things
if you have many of them.

The goal of the API models is to separate the representation the users see from the internal
representation coming from the data/service layers. This way we can update the internal models
in potentially breaking ways without needing to change the models returned to the users for previous
versions of the API, as long as we update those when the internal models change. This becomes much
easier than some other languages because we can take advantage of Go's type checking (and unit tests!)
to alert us at compile time.

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