package request

type route struct {
	middlewares []func(client Client) (bool, int, string)
	request     IRequest
}

func (receiver *route) Middleware(fn func(client Client) (bool, int, string)) *route {
	receiver.middlewares = append(receiver.middlewares, fn)
	return receiver
}

func (receiver *route) Request(request IRequest) *route {
	receiver.request = request
	return receiver
}
