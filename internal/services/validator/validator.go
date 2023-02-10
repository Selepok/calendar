package validator

import (
	"encoding/json"
	"net/http"
)

// Service validates structures
type Service struct {
}

type ok interface {
	OK() error
}

func (s *Service) Validate(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if validatable, ok := v.(ok); ok { // HL
		return validatable.OK() // HL
	} // HL

	return nil
}
