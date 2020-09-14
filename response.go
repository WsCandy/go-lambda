package main

import "encoding/json"

type Response struct {
	Status int              `json:"status"`
	Body   *json.RawMessage `json:"body,omitempty"`
}
