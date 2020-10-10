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

func UpdateData(w http.ResponseWriter, r *http.Request) {
	src.UpdateData(w, r)
}

func SaveChart(w http.ResponseWriter, r *http.Request) {
	src.SaveChart(w, r)
}
