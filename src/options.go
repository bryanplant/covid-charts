package src

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
)

func Options(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)

	locationType := "all"
	if body.LocationType != nil {
		locationType = *body.LocationType
	}

	log.Println("Get Options: " + locationType)

	data, stats := readData(body.LocationType)

	var locations []string
	for _, locationData := range data {
		if locationData.FullName != nil {
			locations = append(locations, *locationData.FullName)
		}
	}

	sort.Strings(locations)
	sort.Strings(stats)

	options := map[string][]string{}
	options["locations"] = locations
	options["stats"] = stats

	bytes, err := json.Marshal(locations)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Options took %s", elapsed)
}
