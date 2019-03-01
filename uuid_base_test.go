package goutils

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/satori/go.uuid"
)

func TestStringUUIDToBase64(t *testing.T) {
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
		result, err := StringUUIDToBase64(test.param)
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

func TestBase64ToUuidString(t *testing.T) {
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
			result:      uuid.Nil.String(),
			resultError: nil,
		},
	}

	for _, test := range tests {
		result, err := Base64ToUuidString(test.param)
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

func TestUUIDToBase64(t *testing.T) {
	tests := []struct {
		param  uuid.UUID
		result string
	}{
		{
			param:  uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
			result: "gx4ABCRlR3+ntBFsrIdXmA==",
		},
		{
			param:  uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
			result: "GvCZS2kBSJa++XXrrFXBug==",
		},
		{
			param:  uuid.FromStringOrNil("bad uuid"),
			result: "AAAAAAAAAAAAAAAAAAAAAA==",
		},
		{
			param:  uuid.FromStringOrNil("1af0994bw6901w4896wbef9w75ebac55c1ba"),
			result: "AAAAAAAAAAAAAAAAAAAAAA==",
		},
	}

	for _, test := range tests {
		result := UUIDToBase64(test.param)
		if !reflect.DeepEqual(test.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, test.result)
		}
	}
}

func TestBase64ToUuid(t *testing.T) {
	tests := []struct {
		param       string
		result      uuid.UUID
		resultError error
	}{
		{
			param:       "gx4ABCRlR3+ntBFsrIdXmA==",
			result:      uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
			resultError: nil,
		},
		{
			param:       "GvCZS2kBSJa++XXrrFXBug==",
			result:      uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
			resultError: nil,
		},
		{
			param:       "AAAAAAAAAAAAAAAAAAAAAA==",
			result:      uuid.Nil,
			resultError: nil,
		},
		{
			param:       "1231244324234",
			result:      uuid.Nil,
			resultError: fmt.Errorf("illegal base64 data at input byte 12"),
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

func TestUUIDToBase58(t *testing.T) {
	tests := []struct {
		param  uuid.UUID
		result string
	}{
		{
			param:  uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
			result: "HC5Dhcue3bEy1mtfgKtnDM",
		},
		{
			param:  uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
			result: "4KwqbCxBYHMKS5EzCkY9sX",
		},
		{
			param:  uuid.FromStringOrNil("bad uuid"),
			result: "1111111111111111",
		},
		{
			param:  uuid.FromStringOrNil("1af0994bw6901w4896wbef9w75ebac55c1ba"),
			result: "1111111111111111",
		},
	}
	for _, tt := range tests {
		result := UUIDToBase58(tt.param)
		if !reflect.DeepEqual(tt.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, tt.result)
		}
	}
}

func TestBase58ToUuid(t *testing.T) {
	tests := []struct {
		param  string
		result uuid.UUID
	}{
		{
			param:  "HC5Dhcue3bEy1mtfgKtnDM",
			result: uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
		},
		{
			param:  "4KwqbCxBYHMKS5EzCkY9sX",
			result: uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
		},
		{
			param:  "1111111111111111",
			result: uuid.Nil,
		},
		{
			param:  "1231244324234",
			result: uuid.Nil,
		},
	}

	for _, test := range tests {
		result := Base58ToUuid(test.param)
		if !reflect.DeepEqual(test.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, test.result)
		}
	}
}

func TestUUIDToBase58Check(t *testing.T) {
	tests := []struct {
		param  uuid.UUID
		result string
	}{
		{
			param:  uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
			result: "3ccy4nXQMJzPVtyCKJw1MKT6jq1tw",
		},
		{
			param:  uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
			result: "3bAnfwuVM9JCEwXF6jqK3d1VSZxQS",
		},
		{
			param:  uuid.FromStringOrNil("bad uuid"),
			result: "3ao273svZEsab9pArFks34VkYnEWh",
		},
		{
			param:  uuid.FromStringOrNil("1af0994bw6901w4896wbef9w75ebac55c1ba"),
			result: "3ao273svZEsab9pArFks34VkYnEWh",
		},
	}
	for _, tt := range tests {
		result := UUIDToBase58Check(tt.param)
		if !reflect.DeepEqual(tt.result, result) {
			t.Errorf("received value %s is not equal to expected value: %s", result, tt.result)
		}
	}
}

func TestBase58CheckToUuid(t *testing.T) {
	tests := []struct {
		param       string
		result      uuid.UUID
		resultError error
	}{
		{
			param:       "3ccy4nXQMJzPVtyCKJw1MKT6jq1tw",
			result:      uuid.FromStringOrNil("831e0004-2465-477f-a7b4-116cac875798"),
			resultError: nil,
		},
		{
			param:       "3bAnfwuVM9JCEwXF6jqK3d1VSZxQS",
			result:      uuid.FromStringOrNil("1af0994b-6901-4896-bef9-75ebac55c1ba"),
			resultError: nil,
		},
		{
			param:       "3ao273svZEsab9pArFks34VkYnEWh",
			result:      uuid.Nil,
			resultError: nil,
		},
		{
			param:       "3Z47KqDctJfEQwXDToadSnzBnofSY",
			result:      uuid.Nil,
			resultError: fmt.Errorf("bad base58 version 41, when expects 42"),
		},
		{
			param:       "1231244324234",
			result:      uuid.Nil,
			resultError: fmt.Errorf("checksum error"),
		},
	}

	for _, test := range tests {
		result, err := Base58CheckToUuid(test.param)
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
