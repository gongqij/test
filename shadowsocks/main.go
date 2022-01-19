package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var logger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags)

func logf(f string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(f, v...))
}

func main() {
	addr := ":20200"
	l, err := net.Listen("tcp", ":20200")
	if err != nil {
		logf("failed to listen on %s: %v", addr, err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			logf("failed to accept: %s", err)
			continue
		}
		go func() {
			defer c.Close()
			c.(*net.TCPConn).SetKeepAlive(true)

			tgt, err := getAddr(c)
			if err != nil {
				logf("failed to get target address: %v", err)
				return
			}

			rc, err := net.Dial("tcp", tgt.String())
			if err != nil {
				logf("failed to connect to target: %v", err)
				return
			}
			defer rc.Close()
			rc.(*net.TCPConn).SetKeepAlive(true)

			logf("proxy %s <-> %s", c.RemoteAddr(), tgt)
			_, _, err = relay(c, rc)
			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					return // ignore i/o timeout
				}
				logf("relay error: %v", err)
			}
		}()
	}
}

func relay(left, right net.Conn) (int64, int64, error) {
	type res struct {
		N   int64
		Err error
	}
	ch := make(chan res)

	go func() {
		n, err := io.Copy(right, left)
		right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
		left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
		ch <- res{n, err}
	}()

	n, err := io.Copy(left, right)
	right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
	left.SetDeadline(time.Now())  // wake up the other goroutine blocking on left
	rs := <-ch

	if err == nil {
		err = rs.Err
	}
	return n, rs.N, err
}
func getAddr(c net.Conn) (Addr, error) {
	return Handshake(c)
}
func relay2(left, right net.Conn) (int64, int64, error) {
	type res struct {
		N   int64
		Err error
	}
	fmt.Println("哈哈哈哈")
	n, err := io.Copy(right, left)
	fmt.Println("嘻嘻嘻嘻", err)
	n, err = io.Copy(left, right)
	fmt.Println("", err)
	return n, 0, err
}
