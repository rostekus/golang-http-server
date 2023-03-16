package socket

import (
	"net"
	"os"
	"syscall"
)

type socketNet struct {
	fd int
}

func NewNetSocket() (*socketNet, error) {
	syscall.ForkLock.Lock()

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, os.NewSyscallError("socket", err)
	}
	syscall.ForkLock.Unlock()

	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		syscall.Close(fd)
		return nil, os.NewSyscallError("setsockopt", err)
	}

	return &socketNet{fd: fd}, nil
}

func (ns *socketNet) Bind(ip net.IP, port int) error {
	sa := &syscall.SockaddrInet4{Port: port}
	copy(sa.Addr[:], ip)
	if err := syscall.Bind(ns.fd, sa); err != nil {
		return os.NewSyscallError("bind", err)
	}
	return nil
}

func (ns *socketNet) Listen() error {
	if err := syscall.Listen(ns.fd, syscall.SOMAXCONN); err != nil {
		return os.NewSyscallError("listen", err)
	}
	return nil
}

func (ns *socketNet) Close() error {
	return syscall.Close(ns.fd)
}

func (ns *socketNet) Accept() (*socketNet, error) {
	// syscall.ForkLock doc states lock not needed for blocking accept.
	nfd, _, err := syscall.Accept(ns.fd)
	if err == nil {
		syscall.CloseOnExec(nfd)
	}
	if err != nil {
		return nil, err
	}
	return &socketNet{nfd}, nil
}

func (ns socketNet) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	n, err := syscall.Read(ns.fd, p)
	if err != nil {
		n = 0
	}
	return n, err
}

func (ns socketNet) Write(p []byte) (int, error) {
	n, err := syscall.Write(ns.fd, p)
	if err != nil {
		n = 0
	}
	return n, err
}
