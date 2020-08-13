package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"golang.org/x/net/context/ctxhttp"
)

func Handler(ctx context.Context) (int, error) {
	ctx, seg := xray.BeginSubsegment(ctx, "MySubsegment")
	err := seg.AddMetadata("url", "https://golang.org/")
	if err != nil {
		return 500, err
	}
	time.Sleep(time.Second * 2)
	httpClient := xray.Client(http.DefaultClient)
	resp, err := ctxhttp.Get(ctx, httpClient, "https://golang.org/")
	seg.Close(err)

	return resp.StatusCode, err
}

func main() {
	lambda.Start(Handler)
}
