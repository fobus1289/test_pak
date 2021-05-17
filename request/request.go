package request

type IRequest interface {
	Construct()
	Authorize(client *Client)
	Rules(validator IValidator)
	FailedValidation(client *Client)
}

type Request struct {
	Entity interface{}
}

func (request *Request) Authorize(client *Client) bool {
	return true
}

func (request *Request) Rules() []string {
	return []string{}
}

func (request *Request) FailedValidation(client *Client) {

}
