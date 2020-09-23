package common

import (
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"net/http"
)

func Serve(fn func(http.ResponseWriter, *http.Request), port string) {
	ctx := context.Background()

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", fn); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
