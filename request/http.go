package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var HandleFunction = map[*regexp.Regexp]*routeMatch{}

type routeMatch struct {
	Url    string
	Method string
	Action func(writer http.ResponseWriter, request *http.Request)
	Route  route
}

func (match *routeMatch) Valid(response http.ResponseWriter, request *http.Request) (ok bool, message string, code int) {

	contend := request.Header.Get("Content-Type")

	if strings.HasPrefix(contend, "multipart/form-data") ||
		strings.HasSuffix(contend, "application/x-www-form-urlencoded") {
		//	fmt.Println(request.PostFormValue("id"))
	} else if strings.HasPrefix(contend, "application/json") {
		data, _ := ioutil.ReadAll(request.Body)
		fmt.Println(string(data))
	} else {
		return false, http.StatusText(415), http.StatusUnsupportedMediaType
	}

	_route := match.Route

	if match.Method != "ANY" && request.Method != match.Method {
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

	var client Client
	rou := new(route)

	var fun = func(writer http.ResponseWriter, request *http.Request) {

		client = Client{
			Request:  request,
			Response: writer,
		}

		action(&client)
	}

	HandleFunction[&_Regexp] = &routeMatch{
		Url:    url,
		Method: method,
		Action: fun,
		Route:  *rou,
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

func Any(url string, action func(client *Client)) *route {
	return handleFunc(url, "ANY", action)
}
