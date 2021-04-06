package main

import (
	"fmt"
	"github.com/fobus1289/test_pak/request"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"time"
)

type Name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func asd(bb bool) {

	if !bb {
		fmt.Println("hello man")
	}

	pc := make([]uintptr, 16)
	n := runtime.Callers(1, pc)
	fmt.Println(pc[:n])
	frames := runtime.CallersFrames(pc[:n])
	frame, more := frames.Next()
	var name = make([]byte, 11111)
	ad := runtime.Stack(name, true)

	fmt.Println(string(name))
	fmt.Println(ad)
	fmt.Println(more)
	fmt.Println(frame)

}

func reg() {
	r, _ := regexp.Compile(`^(/?)\S+(/?)$`)
	result := r.FindString("a2a_66")
	match := r.MatchString("a2a66")
	println(result)
	println(match)
}

var gg = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(r.Header.Get("Content-Type"))
	for key, value := range request.HandleFunction {
		if key.MatchString(r.URL.Path) {
			ok, message, code := value.Valid(w, r)
			if !ok {
				http.Error(w, message, code)
				return
			}
			(*value.Action)(w, r)
			return
		}
	}

	http.NotFound(w, r)
})

func main() {

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
