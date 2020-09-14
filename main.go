package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event Event) (*Response, error) {
	detail, err := event.process()

	if err != nil {
		return nil, err
	}

	return detail.Process(ctx)
}

func main() {
	lambda.Start(HandleRequest)
}
