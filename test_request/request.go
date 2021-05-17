package test_request

import (
	"github.com/fobus1289/test_pak/request"
)

type Request struct {
	message map[string]string
}

func (request *Request) Construct() {

}

func (request *Request) Authorize(client *request.Client) bool {

	return false
}

func (request *Request) Rules(validator request.IValidator) {
	validator.Required("", "").
		Min("", 7, "").
		Max("", 1, "").
		Number("", "")
}

func (request *Request) FailedValidation(client *request.Client) {

}
