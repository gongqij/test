package main

import (
	"io"
	"net"
)

type streamConn struct {
	net.Conn
	r *reader
	w *writer
}

type reader struct {
	io.Reader
}

type writer struct {
	io.Writer
}
