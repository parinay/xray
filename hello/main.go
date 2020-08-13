package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"golang.org/x/net/context/ctxhttp"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	httpClient := xray.Client(http.DefaultClient)
	resp, err := ctxhttp.Get(ctx, httpClient, "https://golang.org/")

	return events.APIGatewayProxyResponse{
		Body:       "Response body",
		StatusCode: resp.StatusCode,
	}, err
}

func main() {
	lambda.Start(Handler)
}
