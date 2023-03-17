package http

import (
	"io"
	"log"
	"net"
	"rostekus/golang-http-server/pkg/mynet/socket"
)

type router interface {
	ServeHTTP(io.Reader, io.Writer) error
}

type server struct {
	addr net.IP
	port int
}

func NewServer(addr string, port int) *server {
	s := server{addr: net.ParseIP(addr),
		port: port}
	return &s
}

func (s *server) ListenAndServe(r router) error {
	socket, err := socket.NewNetSocket()
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	socket.Bind(s.addr, s.port)
	socket.Listen()

	log.Print("================")
	log.Print("Server Started!")
	log.Print("================")
	log.Print()
	log.Printf("addr: http://%s:%d", s.addr, s.port)

	for {

		rw, e := socket.Accept()

		if e != nil {
			panic(e)
		}

		go func() {

			err = r.ServeHTTP(rw, rw)
			if err != nil {
				panic(err)
			}
			rw.Close()
		}()
	}
}
