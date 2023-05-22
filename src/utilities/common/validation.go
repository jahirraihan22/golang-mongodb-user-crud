package common

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var once sync.Once
var LocalValidator *validator.Validate

func init() {
	validate()
}

func validate() *validator.Validate {
	once.Do(func() {
		LocalValidator = validator.New()
	})
	return LocalValidator
}
