package upload

type AvailableTypes struct {
	Types string `json:"type" validate:"oneof=backup current"`
}
