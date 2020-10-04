package src

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
)

func UpdateData(w http.ResponseWriter, _ *http.Request) {
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

	updateOptions(ctx, client, CountryStats, countryStatOptions)
	updateOptions(ctx, client, StateStats, stateStatOptions)

	updateLocations(ctx, client, allLocations)

	elapsed := time.Since(start)
	log.Printf("UpdateData took %s", elapsed)
}

func updateLocationOptions(ctx context.Context, client *firestore.Client, locations map[string]*Location, locationType string) {
	var options []string
	for name := range locations {
		options = append(options, name)
	}

	updateOptions(ctx, client, locationType, options)
}

func updateOptions(ctx context.Context, client *firestore.Client, optionName string, options []string) {
	_, err := client.Collection("options").Doc(optionName).Set(ctx, map[string]interface{}{
		"list": options,
	})
	if err != nil {
		log.Fatalf("Failed adding option: %s, %v", optionName, err)
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

func readCountries() map[string]*Location {
	var countries map[string]*Location
	getDataFromURL(CountryURL, &countries)

	locations := map[string]*Location{}
	for abbreviation, country := range countries {
		if abbreviation == "GEO" {
			// Remove Georgia -- conflicts with state
			continue
		}

		// Parse dates from strings
		for _, record := range country.Data {
			date, err := parseDate(*record.JsonDate)
			if err != nil {
				log.Fatal("Could not parse country record date: " + *record.JsonDate)
			}
			record.Date = date
		}

		country.Type = getStringPointer(LocationTypeCountry)
		country.Abbreviation = getStringPointer(abbreviation)
		country.Color = getStringPointer("")
		country.populateSmoothedData()

		locations[*country.FullName] = country
	}

	return locations
}

func readStates() map[string]*Location {
	var stateRecords []*StateRecord
	getDataFromURL(StateURL, &stateRecords)

	var stateMetadata map[string]StateMetadata
	getDataFromString(StateMetadataJSON, &stateMetadata)

	states := StateRecordsToLocations(stateRecords, stateMetadata)

	locations := map[string]*Location{}
	for name, state := range states {
		if *state.Type == LocationTypeState {
			state.populateSmoothedData()
			locations[name] = state
		}
	}

	return locations
}

func getDataFromURL(url string, data interface{}) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get url: %s %v", url, err)
		return
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func getDataFromString(s string, data interface{}) {
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatal(err)
	}
}

var countryStatOptions = []string {
	TotalCases,
	TotalCasesPerMillion,
	NewCases,
	NewCasesSmoothed,
	NewCasesPerMillion,
	NewCasesSmoothedPerMillion,

	TotalDeaths,
	TotalDeathsPerMillion,
	NewDeaths,
	NewDeathsSmoothed,
	NewDeathsPerMillion,
	NewDeathsSmoothedPerMillion,

	TotalTests,
	TotalTestsPerThousand,
	NewTests,
	NewTestsSmoothed,
	NewTestsPerThousand,
	NewTestsSmoothedPerThousand,

	PositiveRate,
	PositiveRateSmoothed,
}

var stateStatOptions = []string {
	TotalCases,
	TotalCasesPerMillion,
	NewCases,
	NewCasesSmoothed,
	NewCasesPerMillion,
	NewCasesSmoothedPerMillion,

	TotalDeaths,
	TotalDeathsPerMillion,
	NewDeaths,
	NewDeathsSmoothed,
	NewDeathsPerMillion,
	NewDeathsSmoothedPerMillion,

	TotalTests,
	TotalTestsPerThousand,
	NewTests,
	NewTestsSmoothed,
	NewTestsPerThousand,
	NewTestsSmoothedPerThousand,

	PositiveRate,
	PositiveRateSmoothed,
}
