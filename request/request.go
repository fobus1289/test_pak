package request

type IRequest interface {
	Authorize(*Client) bool
	Rules() []string
	FailedValidation(*Client) bool
	ContentTypes() []string
}

type Request struct {
	Entity interface{}
}

func (request *Request) ContentTypes() []string {
	return []string{}
}

func (request *Request) Authorize(client *Client) bool {
	return true
}

func (request *Request) Rules() []string {
	return []string{}
}

func (request *Request) FailedValidation(client *Client) bool {
	return true
}
