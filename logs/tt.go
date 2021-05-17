package logs

type QQQ struct {
	Id   int    `validate:"required,min:6,max:10"`
	Name string `validate:"name,omitempty"`
}
