package main

import (
	"fmt"
	"github.com/fobus1289/test_pak/logger"
	"github.com/fobus1289/test_pak/request"
	"log"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

var Logger = logger.New()

func a(logger2 *logger.Logger, string2 string) {
	logger2.INFO(string2)
}

func b(logger2 *logger.Logger, string2 string) {
	logger2.WARNING(string2)
}

func c(logger2 *logger.Logger, string2 string) {
	logger2.ERROR(string2)
}

var gg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Header.Get("Content-Type"))

	for key, value := range request.HandleFunction {
		if key.MatchString(r.URL.Path) {
			ok, message, code := value.Valid(w, r)
			if !ok {

				go a(Logger, message)
				go b(Logger, message)
				go c(Logger, message)

				http.Error(w, message, code)
				return
			}
			value.Action(w, r)
			return
		}
	}

	http.NotFound(w, r)
})

type Name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func t(ma map[string]string) {
	ma["t"] = "q"
}

func q(ma *map[string]string) {
	(*ma)["Q"] = "Q"
}

type nameq struct {
	Name  *string
	Name1 string
	Bbb   []byte
	ma    map[string]string
}

func qq(n *nameq) *nameq {
	n.Name1 = "111"
	n.Name = &n.Name1
	n.Name1 = strconv.Itoa(time.Now().Nanosecond())
	return n
}

func main() {

	na := &nameq{}
	qq(na)
	ptr := unsafe.Pointer(na)

	na2 := &(*(*nameq)(ptr))

	fmt.Println(na2)
	//na2.Bbb = make([]byte, 10)
	na2.Name1 = "123213"
	na2.Bbb = append(na2.Bbb, 22)
	//na2.Bbb = nil
	fmt.Println(na)

	fmt.Println(na2)
	return
	request.Any("/qq", func(client *request.Client) {
		client.Send("hello man")
	})

	request.Get("/", func(client *request.Client) {
		client.Send("hello man")
	})

	request.Get("/", func(client *request.Client) {
		client.Send("hello man")
	}).Middleware(func(client request.Client) (bool, int, string) {
		return true, 200, "huh"
	}).Request(&request.Request{})

	request.Post("/get/", func(client *request.Client) {
		client.SendJson(Name{
			Id:   1,
			Name: "hello",
		})
	})

	request.Get("g", func(client *request.Client) {
		client.Send("boom")
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        gg,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    20 * time.Second,
	}

	//server := &http.Server{
	//	Addr: ":8888",
	//	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		if r.Method == http.MethodConnect {
	//			handleTunneling(w, r)
	//		} else {
	//			handleHTTP(w, r)
	//		}
	//	}),
	//}
	log.Fatal(s.ListenAndServe())
}

type Service interface {
	Conns()
}

type MainService struct {
	A int
	B int
}

type MainService2 struct {
	A int
	B int
}

func (M MainService2) Conns() {
	println("Conns")
}

func (M MainService) Connsq(service ...Service) {
	service[0].Conns()
}

func (M MainService) Conns() {
	println("Conns")
}

type Controller interface {
	Service(Service)
}

type MainController struct {
	MainService *MainService
}

func (M MainController) Service(service Service) {
	mainService := service.(MainService)
	M.MainService = &mainService
}
