package main

import (
	"encoding/json"
	"errors"
)

type Event struct {
	DetailType DetailType      `json:"detail-type"`
	Detail     json.RawMessage `json:"detail"`
}

type DetailType string

const (
	Example DetailType = "Example"
)

func (event Event) process() (detail Detail, err error) {
	switch event.DetailType {
	case Example:
		detail = new(ExampleDetail)
	default:
		return nil, errors.New("unhandled input")
	}

	if err := detail.Unmarshal(event.Detail); err != nil {
		return nil, err
	}

	return detail, nil
}
