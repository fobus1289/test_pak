package request

import (
	"encoding/json"
	"net/http"
	"unsafe"
)

type Client struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (client *Client) getRequest() *http.Request {
	return client.Request
}

func (client *Client) getResponse() *http.ResponseWriter {
	return &client.Response
}

func (client *Client) Send(message string) {

	client.Response.Header().Set("Content-Type", "text/plain")

	ptr := unsafe.Pointer(&message)

	response := *(*[]byte)(ptr)

	_, err := client.Response.Write(response)

	if err != nil {
		println(err.Error())
	}
}

func (client *Client) SendJson(message interface{}) {

	client.Response.Header().Set("Content-Type", "application/json")

	response, marshalErr := json.Marshal(&message)

	if marshalErr != nil {
		println(marshalErr.Error())
		return
	}

	_, err := client.Response.Write(response)

	if err != nil {
		println(err.Error())
	}
}
