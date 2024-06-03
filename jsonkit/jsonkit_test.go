package jsonkit

import (
	"bytes"
	"strings"
	"testing"
)

type TestType struct {
	Message string `json:"message"`
}

func TestEncode(t *testing.T) {
	jsonToEncode := TestType{
		Message: "hello world",
	}
	json := "{\"message\":\"hello world\"}\n"
	b := new(bytes.Buffer)
	err := Encode(b, jsonToEncode)
	if b.String() != json || err != nil {
		t.Fatalf("error encoding, want = %q, have %q, %v", json, b.String(), err)
	}
}

func TestDecode(t *testing.T) {
	jsonToDecode := "{\"message\":\"hello world\"}\n"
	r := strings.NewReader(jsonToDecode)
	decoded, err := Decode[TestType](r)
	goal := TestType{
		Message: "hello world",
	}
	if decoded != goal || err != nil {
		t.Fatalf("error encoding, want = %q, have %q, %v", goal, decoded, err)
	}
}
