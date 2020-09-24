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

	countries, _ := readData()

	var lines []Line
	for _, country := range body.Countries {
		line := getLine(countries, country, body.YStat)
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
		if locationData.Location != nil {
			locations = append(locations, *locationData.Location)
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
		if item.getField(stat) != nil {
			dataPoint := DataPoint{
				Date:  item.getDate().Format(DateLayout),
				Value: *item.getField(stat),
			}
			dataPoints = append(dataPoints, dataPoint)
		}
	}

	return Line{
		Label: *locationData.Location,
		Data: dataPoints,
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
	var locations map[string]*Location
	getDataFromFile(CountryFile, &locations)

	var stateRecords []StateRecord
	getDataFromFile(StateFile, &stateRecords)

	var stateMetadata map[string]StateMetadata
	getDataFromFile(StateMetadataFile, &stateMetadata)

	states := StateRecordsToLocations(stateRecords, stateMetadata)
	for name, state := range states {
		locations[name] = state
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
	Countries []string `json:"countries"`
	XStat string `json:"x_stat"`
	YStat string `json:"y_stat"`
}

type Line struct {
	Label string
	Data []DataPoint
}

type DataPoint struct {
	Date string
	Value float32
}