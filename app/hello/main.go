package main

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"

    "github.com/ibnr2hc/AWS-Starter-Go/app/hello/types"
)

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context) (Response, error) {
    responseBody := types.OutputEvent{
        Message: "Hello, World!",
    }

    // Convert response body to string
    responseBodyStr, err := json.Marshal(responseBody); if err != nil {
        return Response{StatusCode: 500, Body: ""}, fmt.Errorf("Failed to convert response body to string <- %v", err)
    }
    // Return the response
    return Response{StatusCode: 200, Body: string(responseBodyStr)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
