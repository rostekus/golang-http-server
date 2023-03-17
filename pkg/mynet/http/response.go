package http

import (
	"fmt"
	"net/textproto"
)

type Response struct {
	proto   string // "HTTP/1.1"
	status  int
	message string
	header  textproto.MIMEHeader
	body    string
}

// Create a new Response object
func NewResponse() *Response {
	return &Response{
		proto:   "HTTP/1.1",
		status:  200,
		message: "OK",
		body:    "Hello World!",
		header:  make(textproto.MIMEHeader),
	}
}

func (r *Response) Seriazie() []byte {
	return []byte(r.ToString())
}

func (r *Response) ToString() string {
	responseString := fmt.Sprintf("HTTP/1.1 %d %s\r\n"+
		"Content-Type: text/html; charset=utf-8\r\n"+
		"Content-Length: %d\r\n"+
		"\r\n"+
		"<h1>%s</h1>", r.status, r.message, len(r.body)+9, r.body) // +9 for the <h1></h1> tags
	fmt.Println(responseString)
	return responseString
}
