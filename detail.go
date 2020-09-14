package main

import "context"

type Detail interface {
	Unmarshal(raw []byte) error
	Process(ctx context.Context) (*Response, error)
}
