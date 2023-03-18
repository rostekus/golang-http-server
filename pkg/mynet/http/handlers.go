package http

import (
	"io"
	"os"
)

func NotFoundHandler(r io.Reader, w io.Writer) error {
	_, err := w.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	return err
}

func ErrorHandlerWrapper(status_code int, error_message string, w io.Writer) error {
	resp := NewResponse(status_code, "OK", error_message)
	_, err := w.Write([]byte(resp.ToString()))
	return err
}

func StaticPageHandler(status_code int, filename string, w io.Writer) error {
	// read file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// read file content
	file_content := make([]byte, 1024)
	_, err = file.Read(file_content)
	if err != nil {
		return err
	}
	resp := NewResponse(status_code, "OK", string(file_content))
	w.Write([]byte(resp.ToString()))
	return err
}
