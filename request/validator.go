package request

import "strconv"

type IValidator interface {
	Required(key string, invalid string) IValidator
	Min(key string, length int, invalid string) IValidator
	Max(key string, length int, invalid string) IValidator
	Number(key string, invalid string) IValidator
	Float(key string, invalid string) IValidator
}

type Validator struct {
	*Client
	Errors []string
}

func (validator *Validator) Required(key string, invalid string) IValidator {

	validator.PostFormValue(key)

	return validator
}

func (validator *Validator) Min(key string, length int, invalid string) IValidator {

	value := validator.Request.PostFormValue(key)

	if len(value) >= length {
		validator.Errors = append(validator.Errors, invalid)
	}

	return validator
}

func (validator *Validator) Max(key string, length int, invalid string) IValidator {

	value := validator.Request.PostFormValue(key)

	if len(value) < length {
		validator.Errors = append(validator.Errors, invalid)
	}

	return validator
}

func (validator *Validator) Number(key string, invalid string) IValidator {

	value := validator.Request.PostFormValue(key)

	_, err := strconv.Atoi(value)

	if err != nil {
		validator.Errors = append(validator.Errors, invalid)
	}

	return validator
}

func (validator *Validator) Float(key string, invalid string) IValidator {

	value := validator.Request.PostFormValue(key)

	_, err := strconv.ParseFloat(value, 10)

	if err != nil {
		validator.Errors = append(validator.Errors, invalid)
	}

	return validator
}
