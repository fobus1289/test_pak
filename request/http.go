package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	StatusUnauthorized = []byte(http.StatusText(405))
	StatusBadRequest   = []byte(http.StatusText(400))
	StatusNotFound     = []byte(http.StatusText(404))
)

var HandleFunction = map[*regexp.Regexp]*routeMatch{}

type routeMatch struct {
	Url    string
	Method string
	Action *func(writer http.ResponseWriter, request *http.Request)
	Route  *route
}

func (match *routeMatch) Valid(response http.ResponseWriter, request *http.Request) (ok bool, message string, code int) {

	contend := request.Header.Get("Content-Type")

	if strings.HasPrefix(contend, "multipart/form-data") ||
		strings.HasSuffix(contend, "application/x-www-form-urlencoded") {
		fmt.Println(request.PostFormValue("id"))
	} else if strings.HasPrefix(contend, "application/json") {
		data, _ := ioutil.ReadAll(request.Body)
		fmt.Println(string(data))
	} else {
		return false, http.StatusText(415), http.StatusUnsupportedMediaType
	}

	_route := *match.Route

	if request.Method != match.Method {
		return false, http.StatusText(405), http.StatusMethodNotAllowed
	}

	middlewares := _route.middlewares

	if middlewares != nil {
		for i := 0; i < len(middlewares); i++ {

			if middlewares[i] == nil {
				continue
			}

			_error, code, mgs := middlewares[i](Client{
				Response: response,
				Request:  request,
			})

			if _error {
				return false, mgs, code
			}

		}
	}

	return true, "", 200
}

func handleFunc(url string, method string, action func(client *Client)) *route {

	l := len(url)
	var _Regexp regexp.Regexp

	if l == 1 {
		if strings.HasPrefix(url, "/") || url == "" {
			url = "/"
			_regexp, _ := regexp.Compile(`^/$`)
			_Regexp = *_regexp

		} else {
			tmpUrl := fmt.Sprintf(`^(/?)%s(/?)$`, url)
			_regexp, _ := regexp.Compile(tmpUrl)
			_Regexp = *_regexp
			url = fmt.Sprintf(`/%s/`, url)
		}
	}

	if l != 1 {

		tmpUrl := strings.TrimPrefix(url, "/")
		tmpUrl = strings.TrimSuffix(tmpUrl, "/")
		tmpUrl = fmt.Sprintf(`^(/?)%s(/?)$`, tmpUrl)
		_regexp, _ := regexp.Compile(tmpUrl)
		_Regexp = *_regexp

		if !strings.HasPrefix(url, "/") {
			url = fmt.Sprintf(`/%s`, url)
		}

		if !strings.HasSuffix(url, "/") {
			url = fmt.Sprintf(`%s/`, url)
		}

	}

	var client *Client
	rou := new(route)

	var fun = func(writer http.ResponseWriter, request *http.Request) {

		client = &Client{
			Request:  request,
			Response: writer,
		}

		action(client)
	}

	HandleFunction[&_Regexp] = &routeMatch{
		Url:    url,
		Method: method,
		Action: &fun,
		Route:  rou,
	}

	return rou
}

func Get(url string, action func(client *Client)) *route {
	return handleFunc(url, "GET", action)
}

func Post(url string, action func(client *Client)) *route {
	return handleFunc(url, "POST", action)
}

func Put(url string, action func(client *Client)) *route {
	return handleFunc(url, "PUT", action)
}

func Patch(url string, action func(client *Client)) *route {
	return handleFunc(url, "PATCH", action)
}

func Delete(url string, action func(client *Client)) *route {
	return handleFunc(url, "DELETE", action)
}

func hasUrl(request *http.Request) bool {
	//path := request.URL.Path

	//for _, regex := range regexps {
	//	if regex.MatchString(path) {
	//		return true
	//	}
	//}
	return false
}

func valid(writer http.ResponseWriter, request *http.Request, rou *route, method string, client Client) bool {

	if !hasUrl(request) {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write(StatusNotFound)
		//http.Error(writer, "can't read body", http.StatusBadRequest)
		return false
	}

	if request.Method != method {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = writer.Write(StatusUnauthorized)
		return false
	}
	//_, err := ioutil.ReadAll(request.Body)
	//
	//if err != nil {
	//	log.Printf("Error reading body: %v", err)
	//	http.Error(writer, "can't read body", http.StatusBadRequest)
	//	return false
	//}
	middlewares := rou.middlewares

	if middlewares != nil {
		for i := 0; i < len(middlewares); i++ {

			if middlewares[i] == nil {
				continue
			}

			_error, code, mgs := middlewares[i](client)

			if _error {
				writer.WriteHeader(code)
				_, _ = writer.Write([]byte(mgs))
				return false
			}

		}
	}

	return true
}
