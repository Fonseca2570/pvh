package upload

import (
	"compare/globals"
	"compare/validator"
	"testing"
)

func TestValidate(t *testing.T) {

	validator.Init(&globals.App)

	InvalidType := AvailableTypes{
		Types: "Random",
	}

	ValidType := AvailableTypes{
		Types: "backup",
	}

	err := globals.App.Validator.Struct(InvalidType)
	if err == nil {
		t.Errorf("Type Random should be invalid")
	}

	err = globals.App.Validator.Struct(ValidType)
	if err != nil {
		t.Errorf("Type backup should be valid")
	}

}
