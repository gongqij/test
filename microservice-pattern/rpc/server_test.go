package frpc

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestCountService(t *testing.T) {
	s := NewService(&CountServer{})
	s.run()
}

func TestCountRequest(t *testing.T) {
	c := NewClient("tcp", ":20200")
	var result string
	c.Call("Count", []string{"aa", "1"}, &result)
	log.Println(result)
	c.Call("Count", []string{}, &result)
	log.Println(result)
	c.Call("Count", []string{}, &result)
	log.Println(result)
	c.Call("GetCount", []string{}, &result)
	log.Println(result)
	c.Call("Add", []string{"aa"}, &result)
	log.Println(result)
	c.Call("GetCount", []string{}, &result)
	log.Println(result)
}

func TestSubServer(t *testing.T) {
	s := NewService(new(SubServer))
	s.run()
}

func TestSubRequest(t *testing.T) {
	c := NewClient("tcp", ":20200")
	var result string
	c.Call("Sub", []string{"5", "8"}, &result)
	log.Println(result)
}
