package model

import (
	"encoding/json"
	"errors"
)

// Many2One is dificult to estandarize becouse the second field is custom
type Many2One int64

func (s *Many2One) UnmarshalJSON(data []byte) error {
	var slice []any
	if json.Unmarshal(data, &slice) == nil {
		if len(slice) == 2 {
			if receiveFloat, ok := slice[0].(float64); ok {
				*s = Many2One(int64(receiveFloat))
				return nil
			}
		}
	}
	return errors.New("Could not unmarshall Many2One")
}

func (s *Many2One) MarshalJSON() ([]byte, error) {
	if *s == 0 {
		return nil, nil
	}

	return json.Marshal(int64(*s))
}
