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

    // レスポンスボディを文字列に変換
    responseBodyStr, err := json.Marshal(responseBody); if err != nil {
        return Response{StatusCode: 500, Body: ""}, fmt.Errorf("レスポンスボディの文字列変換に失敗しました <- %v", err)
    }
    // レスポンスを返却する
    return Response{StatusCode: 200, Body: string(responseBodyStr)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
