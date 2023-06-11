# namepoints

This is a small library that enables you to quickly create named endpoint functionality with near effortless documentation for API users.

There is a single public function and a single public interface to use.

```go
NewHandler(handlers []NameHandler) (http.Handler, error)
```
This creates a http.Handler for use in the standard library.

The passed in interface appears below.

```go
type NameHandler interface {
	Name() string
	UpdatesData() bool
	Request() *jsonschema.Schema
	Response() *jsonschema.Schema
	GetResponse(r *http.Request) (interface{}, error)
}
```

The generated http.Handler creates an API with the following conventions:

1. All requests will be POSTs at the root
2. The query param `name` will determine which `NameHandler` will handle the http.Request
3. Passing in an unknown name will return all possible `name`s along with their json request/response schemas.

You can see an example of this in the examples/bookapi folder

https://github.com/stuartaroth/namepoints/tree/master/examples/bookapi
