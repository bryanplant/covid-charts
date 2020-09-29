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

	log.Println("Get Chart Data: " + strings.Join(body.Locations, ","))

	locations, _ := readData(nil)

	selections := map[string]bool{}
	for _, location := range body.Locations {
		selections[location] = true
	}

	_, allStates := selections["allStates"]
	_, allCountries := selections["allCountries"]
	delete(selections, "allStates")
	delete(selections, "allCountries")
	for name, location := range locations {
		if allStates && *location.Type == LocationTypeState {
			selections[name] = true
		}

		if allCountries && *location.Type == LocationTypeCountry {
			selections[name] = true
		}
	}

	var lines []Line
	for location, _ := range selections {
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
}

func Options(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)

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
	addCountries := false
	if locationType == nil || *locationType == LocationTypeCountry {
		addCountries = true
	}
	addStates := true
	if locationType == nil || *locationType == LocationTypeState {
		addStates = true
	}

	locations := map[string]*Location{}

	if addCountries {
		var countries map[string]*Location
		getDataFromFile(CountryFile, &countries)
		for abbreviation, country := range countries {
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
	Locations    []string `json:"locations"`
	LocationType *string  `json:"location_type,omitempty"`
	XStat        *string  `json:"x_stat"`
	YStat        *string  `json:"y_stat"`
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
