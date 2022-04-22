package model

import (
	"encoding/json"
	"reflect"

	odoo "github.com/JohnnyVM/godoo/client"
)

type OdooModel interface {
	Fields() []string
}

func TableName[T OdooModel](table T) string {
	sType := reflect.TypeOf(table)
	field := sType.Field(0)
	return field.Tag.Get("value")
}

func Fields[T OdooModel](table T) []any {
	v := reflect.TypeOf(table)
	out := make([]any, 0)
	for i := 0; i < v.NumField(); i++ {
		if reflect.ValueOf(table).Field(i).CanInterface() {
			out = append(out, v.Field(i).Tag.Get("json"))
		}
	}
	return out
}

func Search[T OdooModel](conn *odoo.Client, args []any, opt map[string]any) ([]int64, error) {
	var table T
	return conn.Search(TableName(table), args, opt)
}

// SearchRead fill a slice of T with the fields specified in opt
// if opt is nil by default get all columns defined for Table
// TODO implement a batch load with search and read
func SearchRead[T OdooModel](conn *odoo.Client, args []any, opt map[string]any) ([]T, error) {
	var table T
	if opt == nil {
		opt = map[string]any{"fields": Fields(table)}
	}

	raw, err := conn.SearchRead(TableName(table), args, opt)
	if err != nil {
		return nil, err
	}

	out := make([]T, len(raw))
	for idx, v := range raw {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &out[idx])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
