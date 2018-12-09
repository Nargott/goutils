package goutils

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestUUIDToBase64(t *testing.T) {
	tests := []struct {
		param       string
		result      string
		resultError error
	}{
		{
			param:       "831e0004-2465-477f-a7b4-116cac875798",
			result:      "gx4ABCRlR3+ntBFsrIdXmA==",
			resultError: nil,
		},
		{
			param:       "1af0994b-6901-4896-bef9-75ebac55c1ba",
			result:      "GvCZS2kBSJa++XXrrFXBug==",
			resultError: nil,
		},
		{
			param:       "bad uuid",
			result:      "",
			resultError: fmt.Errorf("uuid: incorrect UUID length: bad uuid"),
		},
		{
			param:       "1af0994bw6901w4896wbef9w75ebac55c1ba",
			result:      "",
			resultError: fmt.Errorf("uuid: incorrect UUID format 1af0994bw6901w4896wbef9w75ebac55c1ba"),
		},
	}

	for _, test := range tests {
		result, err := UUIDToBase64(test.param)
		if err == nil && test.resultError != nil {
			t.Errorf("unexpected <nil> error was occured")
			continue
		}
		if err != nil && test.resultError == nil {
			t.Errorf("unexpected error was occured: %s", err.Error())
			continue
		}
		if err != nil && test.resultError != nil && !strings.EqualFold(err.Error(), test.resultError.Error()) {
			t.Errorf("unexpected error was occured: %s, when expects: %s", err.Error(), test.resultError.Error())
			continue
		}
		if !reflect.DeepEqual(test.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, test.result)
		}
	}
}

func TestBase64ToUuid(t *testing.T) {
	tests := []struct {
		param       string
		result      string
		resultError error
	}{
		{
			param:       "gx4ABCRlR3+ntBFsrIdXmA==",
			result:      "831e0004-2465-477f-a7b4-116cac875798",
			resultError: nil,
		},
		{
			param:       "GvCZS2kBSJa++XXrrFXBug==",
			result:      "1af0994b-6901-4896-bef9-75ebac55c1ba",
			resultError: nil,
		},
		{
			param:       "bad base64 value",
			result:      "",
			resultError: fmt.Errorf("illegal base64 data at input byte 3"),
		},
		{
			param:       "AAAAAAAAAAAAAAAAAAAAAA==", //clean uuid bytes value
			result:      "",
			resultError: nil,
		},
	}

	for _, test := range tests {
		result, err := Base64ToUuid(test.param)
		if err == nil && test.resultError != nil {
			t.Errorf("unexpected <nil> error was occured")
			continue
		}
		if err != nil && test.resultError == nil {
			t.Errorf("unexpected error was occured: %s", err.Error())
			continue
		}
		if err != nil && test.resultError != nil && !strings.EqualFold(err.Error(), test.resultError.Error()) {
			t.Errorf("unexpected error was occured: %s, when expects: %s", err.Error(), test.resultError.Error())
			continue
		}
		if !reflect.DeepEqual(test.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, test.result)
		}
	}
}
