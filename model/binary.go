package model

import (
	"encoding/json"
	"errors"
)

type Binary []byte

func (s *Binary) UnmarshalJSON(data []byte) error {
	var (
		bin []byte
		b   bool
	)
	if json.Unmarshal(data, &bin) == nil {
		*s = Binary(bin)
		return nil
	}
	if json.Unmarshal(data, &b) == nil && b == false {
		*s = nil // or whatever you want
		return nil
	}
	return errors.New("Could not unmarshall OdooString")
}
