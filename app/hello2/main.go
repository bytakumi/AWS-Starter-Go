package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

    "github.com/ibnr2hc/AWS-Starter-Go/app/hello/types"
)

func HandleRequest(ctx context.Context) (types.OutputEvent, error) {
	return types.OutputEvent{Message: "Hello World"}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
