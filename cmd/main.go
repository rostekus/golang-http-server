package main

import (
	"flag"
	"io"
	"rostekus/golang-http-server/pkg/mynet/http"
)

func main() {
	ipFlag := flag.String("ip_addr", "127.0.0.1", "The IP address to use")
	portFlag := flag.Int("port", 8080, "The port to use.")
	flag.Parse()
	port := *portFlag

	srv := http.NewServer(*ipFlag, port)
	router := http.NewRouter()
	router.GET("/", func(rq io.Reader, w io.Writer) error {
		resp := http.NewResponse(200, "OK", "Hello World :)")
		w.Write(resp.Seriazie())
		return nil
	})

	router.GET("/test", func(rq io.Reader, w io.Writer) error {
		err := http.StaticPageHandler(200, "static/index.html", w)
		if err != nil {
			return err
		}
		return nil
	})
	srv.ListenAndServe(router)

}
