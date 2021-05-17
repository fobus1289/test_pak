package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type Client struct {
	Response      http.ResponseWriter
	Request       *http.Request
	GetHeader     func(key string)
	SetHeader     func(key, value string)
	PostFormValue func(key string)
	body          io.ReadCloser
	BodyString    string
	BodyBytes     []byte
}

func (client *Client) ParseBody() {
	body, err := ioutil.ReadAll(client.body)

	if err != nil {
		println(err.Error())
	}

	client.BodyBytes = body
	ptr := unsafe.Pointer(&client.BodyBytes)
	client.BodyString = *(*string)(ptr)

	client.BodyString = "sad"
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
