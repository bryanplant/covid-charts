package main

import (
	"github.com/bryanplant/covid-charts/server/common"
	"github.com/bryanplant/covid-charts/src"
)

func main() {
	common.Serve(src.Options, "8081")
}
