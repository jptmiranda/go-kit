package jsonkit

import (
	"encoding/json"
	"fmt"
	"io"
)

func Encode(w io.Writer, value any) error {
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		return fmt.Errorf("error encoding json: %v", err)
	}
	return nil
}

func Decode[T any](r io.Reader) (T, error) {
	var value T
	if err := json.NewDecoder(r).Decode(&value); err != nil {
		return value, fmt.Errorf("error decoding json: %v", err)
	}
	return value, nil
}
