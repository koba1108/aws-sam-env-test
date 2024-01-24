package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/koba1108/aws-sam-env-test/internal/handler"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	e := echo.New()
	api2 := e.Group("/api2")
	{
		api2.GET("", handler.HealthCheckHandler)
	}
	echoLambda = echoadapter.New(e)
}

func main() {
	lambda.Start(proxy)
}

func proxy(c context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(c, req)
}
