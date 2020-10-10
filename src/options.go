package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
)

func Options(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Get Options")

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	countryOptions := getOptions(ctx, client, LocationTypeCountry)
	stateOptions := getOptions(ctx, client, LocationTypeState)
	countryStats := getOptions(ctx, client, CountryStats)
	stateStats := getOptions(ctx, client, StateStats)

	options := map[string][]string{
		LocationTypeCountry: countryOptions,
		LocationTypeState:   stateOptions,
		CountryStats:        countryStats,
		StateStats:          stateStats,
	}

	bytes, err := json.Marshal(options)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	log.Printf("Get Options took %s", elapsed)
}

func getOptions(ctx context.Context, client *firestore.Client, optionName string) []string {
	var optionsMap map[string][]string
	doc, err := client.Collection("options").Doc(optionName).Get(ctx)
	if err != nil {
		panic(fmt.Sprintf("Could not get options: %s", optionName))
	}
	err = doc.DataTo(&optionsMap)
	if err != nil {
		panic(fmt.Sprintf("Could not cast options to list: %s", optionName))
	}

	options := optionsMap["list"]
	sort.Strings(options)
	return options
}
