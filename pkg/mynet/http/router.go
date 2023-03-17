package http

import (
	"io"
	"log"
	"regexp"
)

type Handler func(io.Reader, io.Writer) error

type routeEntry struct {
	path    *regexp.Regexp
	method  string
	handler Handler
}

type Router struct {
	routes []routeEntry
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) POST(path string, handler Handler) {
	exactPath := regexp.MustCompile("^" + path + "$")
	e := routeEntry{
		path:    exactPath,
		method:  "POST",
		handler: handler}
	r.routes = append(r.routes, e)
}
func (r *Router) GET(path string, handler Handler) {
	exactPath := regexp.MustCompile("^" + path + "$")
	e := routeEntry{
		path:    exactPath,
		method:  "GET",
		handler: handler}
	r.routes = append(r.routes, e)
}

func (r *Router) ServeHTTP(rq io.Reader, w io.Writer) error {
	request, err := ParseRequest(rq)
	if err != nil {
		return err
	}
	log.Printf("request: %s %s", request.method, request.uri)
	for _, e := range r.routes {
		if e.method == request.method && e.path.MatchString(request.uri) {
			err = e.handler(rq, w)
			return err
		}
	}
	return NotFoundHandler(rq, w)
}
