package src

import (
	"context"
	"encoding/json"
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

	body := readRequest(r)

	locationType := LocationTypeAll
	if body.LocationType != nil {
		locationType = *body.LocationType
	}

	log.Println("Get Options: " + locationType)

	ctx := context.Background()
	client := getFirebaseClient(ctx)
	options := getLocationOptions(ctx, client, locationType)

	bytes, err := json.Marshal(options)
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

func getLocationOptions(ctx context.Context, client *firestore.Client, locationType string) []string {
	var optionsMap map[string][]string
	doc, err := client.Collection("options").Doc(locationType).Get(ctx)
	if err != nil {
		log.Fatalf("Could not get options: %s", locationType)
	}
	err = doc.DataTo(&optionsMap)
	if err != nil {
		log.Fatalf("Could not cast options to list: %s", locationType)
	}

	options := optionsMap["list"]
	sort.Strings(options)
	return options
}
