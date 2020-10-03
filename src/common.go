package src

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func getFirebaseClient(ctx context.Context) *firestore.Client {
	// Use the application default credentials
	conf := &firebase.Config{ProjectID: "covid-charts-289323"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func readRequest(r *http.Request) *Request {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := &Request{}
	err = json.Unmarshal([]byte(buf.String()), body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func readData(locationType string) (map[string]*Location, []string) {
	start := time.Now()

	addCountries := locationType == LocationTypeCountry || locationType == LocationTypeAll
	addStates := locationType == LocationTypeState || locationType == LocationTypeAll

	locations := map[string]*Location{}

	if addCountries {
		countries := readCountries()
		for name, country := range countries {
			locations[name] = country
		}
	}

	if addStates {
		states := readStates()
		for name, state := range states {
			locations[name] = state
		}
	}

	elapsed := time.Since(start)
	log.Printf("Reading data took %s", elapsed)

	return locations, []string{}
}

func readCountries() map[string]*Location {
	var countries map[string]*Location
	getDataFromFile(CountryFile, &countries)

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
	getDataFromFile(StateFile, &stateRecords)

	var stateMetadata map[string]StateMetadata
	getDataFromFile(StateMetadataFile, &stateMetadata)

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

func getDataFromFile(filename string, data interface{}) {
	file := getFile(filename)
	reader := io.Reader(file)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func getFile(path string) *os.File {
	var file *os.File
	var err error

	file, err = os.Open(path)
	if err != nil {
		file, err = os.Open("src/" + path)
		if err != nil {
			file, err = os.Open("serverless_function_source_code/src/" + path)
			if err != nil {
				log.Fatalln("Couldn't open the file", err)
			}
		}
	}

	return file
}

type Request struct {
	Locations    []string `json:"locations"`
	LocationType *string  `json:"location_type,omitempty"`
	XStat        *string  `json:"x_stat"`
	YStat        *string  `json:"y_stat"`
}
