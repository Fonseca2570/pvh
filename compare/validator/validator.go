package validator

import (
	"compare/globals"
	"github.com/go-playground/validator/v10"
)

func Init(App *globals.Application) {
	App.Validator = validator.New()
}
