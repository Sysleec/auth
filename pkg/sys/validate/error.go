package validate

import (
	"encoding/json"
	"errors"
)

type ValidationErrors struct {
	Message []string `json:"error_message"`
}

func (v *ValidationErrors) addError(message string) {
	v.Message = append(v.Message, message)
}

func NewValidationError(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Message: messages,
	}
}

func (v *ValidationErrors) Error() string {
	data, err := json.Marshal(v.Message)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}
