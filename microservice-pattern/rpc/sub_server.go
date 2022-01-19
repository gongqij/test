//sub_server.go
package frpc

type SubServer struct {
}

func (s *SubServer) Sub(a, b int) int {
	return a - b
}
