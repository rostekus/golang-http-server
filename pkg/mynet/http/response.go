package http

import (
	"fmt"
	"net/textproto"
	"strconv"
)

type Response struct {
	proto   string // "HTTP/1.1"
	status  int
	message string
	header  textproto.MIMEHeader
	body    string
}

// Create a new Response object
func NewResponse(status int, message string, body string) *Response {
	return &Response{
		proto:   "HTTP/1.1",
		status:  status,
		message: message,
		body:    body,
		header:  make(textproto.MIMEHeader),
	}
}

func (r *Response) SetHeader(key string, value string) {
	r.header.Set(key, value)
}

func (r *Response) Seriazie() []byte {
	return []byte(r.ToString())
}

func (r *Response) ToString() string {
	r.header.Set("Content-Type", "text/html; charset=utf-8")
	r.header.Set("Content-Length", strconv.Itoa(len(r.body))) // +9 for the <h1></h1> tags
	fmt.Println(r.header)
	responseString := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.status, r.message)
	for key, values := range r.header {
		for _, value := range values {
			responseString += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}
	responseString += fmt.Sprintf("\r\n%s", r.body)
	return responseString
}
