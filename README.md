# Example logging middleware in Go

In this example we create a simple middleware logger using Go's `net/http` package only.

We log:
  * Requests with log messages prefixed with `\[Request\] `
  * Responses with log messages prefixed with `\[Response\] `

When logging responses, in addition to reporting the response sent to the client, we also log the response status code and the time it took to process the response.

To achieve all this we wrap Go's `http.ResponseWriter` to capture the status and the response as a buffer. Without wrapping the interface we could not capture the output written to the client or the status code and hence could not log them.

The `httplog` package exposes two public function:
  * `Mux`
  * `LogMux`

`Mux` is the default server mux. It automatically logs to standard out.
`LogMux` allows you to set alternative `log.Logger`s for the request and response logs.

Also provided is a tiny example in `main.go`
