package main

import (
	"context"
	"encoding/json"
)

type ExampleDetail struct {
	ID string `json:"id"`
}

var _ Detail = (*ExampleDetail)(nil)

func (detail *ExampleDetail) Unmarshal(raw []byte) error {
	return json.Unmarshal(raw, &detail)
}

func (detail *ExampleDetail) Process(ctx context.Context) (*Response, error) {
	return &Response{Status: 200}, nil
}
