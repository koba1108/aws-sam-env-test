package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type TestResponse struct {
	TestEnv1 string `json:"testEnv1"`
	TestEnv2 string `json:"testEnv2"`
	TestEnv3 string `json:"testEnv3"`
}

func handler(c context.Context, req *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res := TestResponse{
		TestEnv1: os.Getenv("TEST_ENV_1"),
		TestEnv2: os.Getenv("TEST_ENV_2"),
		TestEnv3: os.Getenv("TEST_ENV_3"),
	}
	j, err := json.Marshal(&res)
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(j),
	}, nil
}

func main() {
	lambda.Start(handler)
}
