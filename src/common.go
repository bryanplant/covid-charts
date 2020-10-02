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

func readData(locationType *string) (map[string]*Location, []string) {
	start := time.Now()

	addCountries := false
	if locationType == nil || *locationType == LocationTypeCountry {
		addCountries = true
	}
	addStates := false
	if locationType == nil || *locationType == LocationTypeState {
		addStates = true
	}

	locations := map[string]*Location{}

	if addCountries {
		var countries map[string]*Location
		getDataFromFile(CountryFile, &countries)
		for abbreviation, country := range countries {
			if abbreviation == "GEO" {
				// Remove Georgia -- conflicts with state
				continue
			}
			country.Type = getStringPointer(LocationTypeCountry)
			country.Abbreviation = getStringPointer(abbreviation)
			country.Color = getStringPointer("")

			locations[*country.FullName] = country
		}
	}

	if addStates {
		var stateRecords []StateRecord
		getDataFromFile(StateFile, &stateRecords)

		var stateMetadata map[string]StateMetadata
		getDataFromFile(StateMetadataFile, &stateMetadata)

		states := StateRecordsToLocations(stateRecords, stateMetadata)
		for name, state := range states {
			if *state.Type == LocationTypeState {
				locations[name] = state
			}
		}
	}

	for _, location := range locations {
		location.populateSmoothedData()
	}

	elapsed := time.Since(start)
	log.Printf("Reading data took %s", elapsed)

	return locations, []string{}
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
