package main

import (
    "github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, name myEvent) (string, error) {
    return "Hello", nil
}

func main() {
    lambda.Start(HandleRequest)
}
