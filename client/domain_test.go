package client

import (
	"fmt"
	"reflect"
	"testing"
)

// DomainEqual check both domains are equal
func domainEqual(l Domain, r Domain) bool {
	for idx, vr := range r {
		switch v := reflect.ValueOf(vr); v.Kind() {
		case reflect.Slice:
			dom, ok := l[idx].(Domain)
			if !ok {
				return false
			}
			return domainEqual(dom, vr.(Domain))
		case reflect.String:
			cad, ok := l[idx].(string)
			if !ok {
				return false
			}
			if vr.(string) != cad {
				return false
			}
		default:
			panic(fmt.Sprintf("Domain contain %v", v))
		}
	}
	return true
}

type domainTest struct {
	arg      Domain
	expected Domain
}

var equalTests = []domainTest{
	{Domain{"column", "op", "value"}, Domain{"column", "op", "value"}},
}

var diferentTests = []domainTest{
	{Domain{"column", "op", "value"}, Domain{"column", "op", "valua"}},
}

func TestEqual(t *testing.T) {
	for idx, test := range equalTests {
		if !domainEqual(equalTests[idx].arg, test.expected) {
			t.Errorf("Output %q not equal to expected %q", equalTests[idx].arg, test.expected)
		}
	}
	for idx, test := range diferentTests {
		if domainEqual(equalTests[idx].arg, test.expected) {
			t.Errorf("Output %q and %q equal", equalTests[idx].arg, test.expected)
		}
	}
}

var andTests = []domainTest{
	{Domain{"column", "op", "value"}, Domain{Domain{"column", "op", "value"}}},
	{Domain{"column", "op", "value"}, Domain{"&", Domain{"column", "op", "value"}, Domain{"column", "op", "value"}}},
}

func TestAnd(t *testing.T) {
	var d Domain
	for _, test := range andTests {
		if output := d.And(test.arg); !domainEqual(*output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

var orTests = []domainTest{
	{Domain{"column", "op", "value"}, Domain{Domain{"column", "op", "value"}}},
	{Domain{"column", "op", "value"}, Domain{"|", Domain{"column", "op", "value"}, Domain{"column", "op", "value"}}},
}

func TestOr(t *testing.T) {
	var d Domain
	for _, test := range orTests {
		if output := d.Or(test.arg); !domainEqual(*output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
