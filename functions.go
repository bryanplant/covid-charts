package functions

import (
	"github.com/bryanplant/covid-charts/src"
	"net/http"
)

func ChartData(w http.ResponseWriter, r *http.Request) {
	src.ChartData(w, r)
}

func Options(w http.ResponseWriter, r *http.Request) {
	src.Options(w, r)
}
