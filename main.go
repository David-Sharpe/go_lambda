package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
    Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    return "Hello", nil
}

func main() {
    lambda.Start(HandleRequest)
}
