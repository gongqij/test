//frpc.go
package frpc

import (
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type Service struct {
	value reflect.Value
}

type Client struct {
	conn net.Conn
}

func NewClient(network, address string) *Client {
	c, err := net.Dial("tcp", ":20200")
	if err != nil {
		log.Error(err)
	}
	return &Client{conn: c}
}

type Result struct {
	Res reflect.Value
	Err error
}

func NewService(caller interface{}) *Service {
	value := reflect.ValueOf(caller)
	log.Printf(value.String())
	return &Service{value: value}
}

func (a Service) run() {
	l, err := net.Listen("tcp", ":20200")
	if err != nil {
		log.Error(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Error(err)
		}
		log.Println(c.RemoteAddr(), c.LocalAddr())

		scan := bufio.NewScanner(c)
		scan.Split(bufio.ScanLines)
		for scan.Scan() {
			str := scan.Text()
			log.Println(str)
			args := strings.Split(str, " ")
			var result Result
			if len(args) > 1 {
				log.Printf("func : %s, args : %v\n", args[0], args[1:])
				result = a.Exec(args[0], args[1:])
			} else {
				log.Printf("func: %s\n", args[0])
				result = a.Exec(args[0], []string{})
			}
			var res = "ok"
			if result.Res.IsValid() {
				switch result.Res.Kind() {
				case reflect.Int:
					i := result.Res.Int()
					res = strconv.Itoa(int(i))
				default:
					res = "null"
				}
			}
			_, err := io.WriteString(c, res)
			if err != nil {
				log.Error(err)
			}
		}
		c.Close()
	}
}

func (a Service) Exec(funcName string, parameters []string) Result {
	f := a.value.MethodByName(funcName)
	if !f.IsValid() {
		return Result{Err: errors.New("func not found")}
	}
	var paramList []reflect.Value

	for _, parameter := range parameters {
		paramList = append(paramList, reflect.ValueOf(parameter))
	}

	t := f.Type()
	for i := 0; i < t.NumIn(); i++ {
		if t.In(i).Kind() == reflect.Int {
			s, _ := strconv.Atoi(parameters[i])
			paramList[i] = reflect.ValueOf(s)
		}
	}
	results := f.Call(paramList)
	if len(results) == 1 {
		log.Println(results[0].Int())
		return Result{Res: results[0]}
	}
	return Result{}
}

func (c Client) Call(funcName string, parameters []string, result *string) {
	request := funcName
	if len(parameters) > 0 {
		request += " " + strings.Join(parameters, " ")
	}
	request += "\n"
	_, err := c.conn.Write([]byte(request))
	if err != nil {
		log.Error(err)
	}
	data := make([]byte, 200)
	n, err := c.conn.Read(data)
	*result = string(data[:n])
}
