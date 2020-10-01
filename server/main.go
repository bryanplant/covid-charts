package main

import (
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/bryanplant/covid-charts/src"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/chart-data", src.ChartData); err != nil {
		panic(err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/options", src.Options); err != nil {
		panic(err)
	}
	if err := funcframework.Start("8080"); err != nil {
		panic(err)
	}
}
