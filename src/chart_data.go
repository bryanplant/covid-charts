package src

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
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

	//ctx := context.Background()
	//client := getFirebaseClient(ctx)
	//_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
	//	"first": "Ada",
	//	"last":  "Lovelace",
	//	"born":  1815,
	//})
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}
	//
	//iter := client.Collection("users").Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}

	body := readRequest(r)

	log.Println("Get Chart Data: " + strings.Join(body.Locations, ","))

	locations, _ := readData(nil)

	selections := map[string]bool{}
	for _, location := range body.Locations {
		selections[location] = true
	}

	_, allStates := selections["All States"]
	_, allCountries := selections["All Countries"]
	delete(selections, "All States")
	delete(selections, "All Countries")
	for name, location := range locations {
		if allStates && *location.Type == LocationTypeState {
			selections[name] = true
		}

		if allCountries && *location.Type == LocationTypeCountry {
			selections[name] = true
		}
	}

	var lines []Line
	for location := range selections {
		locationData, ok := locations[location]
		if !ok {
			log.Println("Could not find location: " + location)
			continue
		}
		line := getLine(locationData, *body.YStat)
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].DisplayName < lines[j].DisplayName
	})

	bytes, err := json.Marshal(lines)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("ChartData took %s", elapsed)
}

func getLine(location *Location, stat string) Line {
	var dataPoints []DataPoint
	for _, item := range location.Data {
		if value := item.getField(stat); value != nil && *value >= 0 {
			dataPoint := DataPoint{
				Date:  item.getDate().Format(DateLayout),
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
