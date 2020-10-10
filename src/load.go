package src

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func LoadChart(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Load Chart")

	body := readRequest(r)

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	doc, err := client.Collection("charts").Doc(*body.ChartID).Get(ctx)
	if err != nil {
		panic("Failed to get chart: " + *body.ChartID)
	}

	var settings ChartSettings
	err = doc.DataTo(&settings)
	if err != nil {
		panic("Failed to Unmarshal chart: " + *body.ChartID)
	}

	bytes, err := json.Marshal(settings)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	log.Printf("Load Chart took %s", elapsed)
}
