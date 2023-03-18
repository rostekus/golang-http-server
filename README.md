# Golang HTTP Server from Scratch

This repository contains an HTTP server implemented from scratch using Linux sockets and the Go programming language.

## Dependencies

To build and run the HTTP server, you will need the following dependencies:

- Go 1.16 or higher
- A Unix-like operating system (tested on Ubuntu 20.04)

## Example

```golang
package main

import (
	"flag"
	"io"
	"rostekus/golang-http-server/pkg/mynet/http"
)

func main() {
	// Parse command-line flags for IP address and port number
	ipFlag := flag.String("ip_addr", "127.0.0.1", "The IP address to use")
	portFlag := flag.Int("port", 8080, "The port to use.")
	flag.Parse()
	port := *portFlag

	// Create a new HTTP server using the specified IP address and port
	srv := http.NewServer(*ipFlag, port)

	// Create a new router to handle incoming requests
	router := http.NewRouter()

	// Define a route for the root URL ("/") that returns a simple "Hello World" message
	router.GET("/", func(rq io.Reader, w io.Writer) error {
		resp := http.NewResponse(200, "OK", "Hello World :)")
		w.Write(resp.Serialize())
		return nil
	})

	// Define a route for the "/test" URL that serves a static HTML page
	router.GET("/test", func(rq io.Reader, w io.Writer) error {
		err := http.StaticPageHandler(200, "static/index.html", w)
		if err != nil {
			return err
		}
		return nil
	})

	// Start the server and listen for incoming requests
	srv.ListenAndServe(router)
}
```

## Building the HTTP server

To build the HTTP server, clone this repository and run the following command:

```bash
go build . 
```

This will generate an executable file named `main`.

## Running the HTTP server

To run the HTTP server, execute the `http-server` executable:

```bash
./main --ip_addr=127.0.0.1 --port=8080
```

The server will start listening for incoming requests on port 8080. You can change the port number by setting the `PORT` environment variable before running the server:


## Testing the HTTP server

To test the HTTP server, you can use any web browser or HTTP client. For example, you can open a web browser and go to `http://localhost:8080/` to see the server's response.

You can also use `curl` to send HTTP requests from the command line:

```bash
curl -v http://localhost:8080/
```
## License

This project is licensed under the MIT License. See the LICENSE file for details.
