package main

import (
	"context"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	functions "github.com/bryanplant/covid-charts"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/chart-data", functions.ChartData); err != nil {
		panic(err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/options", functions.Options); err != nil {
		panic(err)
	}
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/update-data", functions.UpdateData); err != nil {
		panic(err)
	}
	if err := funcframework.Start("8080"); err != nil {
		panic(err)
	}
}
