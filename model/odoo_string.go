package model

import (
	"encoding/json"
	"errors"
)

type OdooString string

func (s *OdooString) UnmarshalJSON(data []byte) error {
	var (
		str string
		b   bool
	)
	if json.Unmarshal(data, &str) == nil {
		*s = OdooString(str)
		return nil
	}
	if json.Unmarshal(data, &b) == nil && b == false {
		*s = OdooString("") // or whatever you want
		return nil
	}
	return errors.New("Could not unmarshall OdooString")
}
