package src

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"net/http"
	"time"
)

func UpdateData(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Update Data")

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	// Update state options
	states := readStates()
	updateLocationOptions(ctx, client, states, LocationTypeState)

	// Update country options
	countries := readCountries()
	updateLocationOptions(ctx, client, countries, LocationTypeCountry)

	// Update all options
	allLocations := map[string]*Location{}
	for k, v := range states {
		allLocations[k] = v
	}
	for k, v := range countries {
		allLocations[k] = v
	}
	updateLocationOptions(ctx, client, allLocations, LocationTypeAll)

	updateLocations(ctx, client, allLocations)

	elapsed := time.Since(start)
	log.Printf("UpdateData took %s", elapsed)
}

func updateLocationOptions(ctx context.Context, client *firestore.Client, locations map[string]*Location, locationType string) {
	var options []string
	for name := range locations {
		options = append(options, name)
	}

	_, err := client.Collection("options").Doc(locationType).Set(ctx, map[string]interface{}{
		"list": options,
	})
	if err != nil {
		log.Fatalf("Failed adding option: %s, %v", locationType, err)
	}
}

func updateLocations(ctx context.Context, client *firestore.Client, locations map[string]*Location) {
	for name, location := range locations {
		_, err := client.Collection("data").Doc(name).Set(ctx, map[string]interface{}{
			"data": location,
		})

		if err != nil {
			log.Fatalf("Failed adding data: %s, %v", "United States", err)
		}
	}
}
