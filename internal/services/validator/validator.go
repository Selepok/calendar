package validator

import (
	"encoding/json"
	"io"
)

// Service validates structures
type Service struct {
}

type Ok interface {
	OK() error
}

func (s *Service) Validate(r io.Reader, v Ok) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	return v.OK()
}
