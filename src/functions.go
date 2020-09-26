package src

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func ChartData(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)

	locations, _ := readData()

	selections := body.Locations
	if len(selections) == 1 && selections[0] == "allStates" || selections[0] == "allCountries" {
		var locationType string
		if selections[0] == "allStates" {
			locationType = LocationTypeState
		} else if selections[0] == "allCountries" {
			locationType = LocationTypeCountry
		}

		selections = []string{}

		for name, location := range locations {
			if location.Type == locationType {
				selections = append(selections, name)
			}
		}
	}

	var lines []Line
	for _, location := range selections {
		line := getLine(locations, location, body.YStat)
		lines = append(lines, line)
	}

	bytes, err := json.Marshal(lines)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func Options(w http.ResponseWriter, _ *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, stats := readData()

	var locations []string
	for _, locationData := range data {
		if locationData.Name != nil {
			locations = append(locations, *locationData.Name)
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
}

func getLine(locations map[string]*Location, location, stat string) Line {
	locationData := locations[location]

	var dataPoints []DataPoint
	for _, item := range locationData.Data {
		if value := item.getField(stat); value != nil && *value >= 0 {
			dataPoint := DataPoint{
				Date:  item.getDate().Format(DateLayout),
				Value: *value,
			}
			dataPoints = append(dataPoints, dataPoint)
		}
	}

	return Line{
		ID:          location,
		DisplayName: *locationData.Name,
		Color:       *locationData.Color,
		Data:        dataPoints,
	}
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

func readData() (map[string]*Location, []string) {
	locations := map[string]*Location{}

	var countries map[string]*Location
	getDataFromFile(CountryFile, &countries)
	for _, country := range countries {
		country.Type = LocationTypeCountry
	}
	for name, country := range countries {
		locations[name] = country
		locations[name].Color = getStringPointer("")
	}

	var stateRecords []StateRecord
	getDataFromFile(StateFile, &stateRecords)

	var stateMetadata map[string]StateMetadata
	getDataFromFile(StateMetadataFile, &stateMetadata)

	states := StateRecordsToLocations(stateRecords, stateMetadata)
	for name, state := range states {
		locations[name] = state
	}

	for _, location := range locations {
		location.populateSmoothedData()
	}

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
			file, err = os.Open("src/src/" + path)
			if err != nil {
				log.Fatalln("Couldn't open the file", err)
			}
		}
	}

	return file
}

type Request struct {
	Locations []string `json:"locations"`
	XStat     string   `json:"x_stat"`
	YStat     string   `json:"y_stat"`
}

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
