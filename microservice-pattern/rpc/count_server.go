// count_server.go
package frpc

type CountServer struct {
	count int
}

func (a *CountServer) GetCount() int {
	return a.count
}

func (a *CountServer) Count() {
	a.count++
}

func (a *CountServer) Add(i int) {
	a.count += i
}
