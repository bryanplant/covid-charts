package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

type Line struct {
	ID          string
	DisplayName string
	Color       string
	Data        []DataPoint
}

type DataPoint struct {
	Date  string
	Value float64
}

func ChartData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)

	log.Println("Get Chart Data: " + strings.Join(body.Locations, ", "))

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	selections := map[string]bool{}
	for _, location := range body.Locations {
		selections[location] = true
	}

	var lines []Line
	for selection := range selections {
		location := getLocation(ctx, client, selection)
		line := getLine(location, *body.YStat)
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].DisplayName < lines[j].DisplayName
	})

	bytes, err := json.Marshal(lines)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	log.Printf("ChartData took %s", elapsed)
}

func getLocation(ctx context.Context, client *firestore.Client, name string) *Location {
	doc, err := client.Collection("data").Doc(name).Get(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed getting location: %s, %v", name, err))
	}

	var locationData map[string]*Location
	err = doc.DataTo(&locationData)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal firestore data: %s, %v", name, err))
	}

	return locationData["data"]
}

func getLine(location *Location, stat string) Line {
	var dataPoints []DataPoint
	for _, item := range location.Data {
		if value := item.getField(stat); value != nil && *value >= 0 {
			dataPoint := DataPoint{
				Date:  item.Date.Format(DateLayout),
				Value: *value,
			}
			dataPoints = append(dataPoints, dataPoint)
		}
	}

	return Line{
		ID:          *location.Abbreviation,
		DisplayName: *location.FullName,
		Color:       *location.Color,
		Data:        dataPoints,
	}
}
