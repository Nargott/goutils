package goutils

import (
    "testing"
    "reflect"
)

func TestHashSHA256(t *testing.T) {
    tests := []struct {
        value string
        expected string
    }{
        {"DataToBeSHA256Hashed", "r0vA_A5S6RVmVJDRRj3aQA9ko8MP_G3TaRgndTNhnZE="},
        {"AnotherDataToBeSHA256Hashed", "fJC9eCLhphPWavkQgRd0nJ60KbMQKWetnak-6WMhH9I="},
    }

    for _, test := range tests {
        result := HashSHA256Base64(test.value)
        if !reflect.DeepEqual(test.expected, result) {
            t.Errorf("received value %s is not equal to expected value: %s", result, test.expected)
        }
    }
}