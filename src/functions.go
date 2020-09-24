package src

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"encoding/csv"
	"encoding/json"
	"net/http"
)

func ChartData(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)

	data, _ := readData()

	var lines []line
	for _, country := range body.Countries {
		line := getLine(data, country, body.YStat)
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

func getLine(locations map[string]location, location, stat string) line {
	locationData := locations[location]

	var dataPoints []dataPoint
	for _, item := range locationData.Data {
		if item.getField(stat) != nil {
			dataPoint := dataPoint{
				Date:  time.Time(item.Date).Format(DateLayout),
				Value: *item.getField(stat),
			}
			dataPoints = append(dataPoints, dataPoint)
		}
	}

	return line{
		Label: *locationData.Location,
		Data: dataPoints,
	}
}

func readRequest(r *http.Request) *request {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := &request{}
	err = json.Unmarshal([]byte(buf.String()), body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func readData() (map[string]location, []string) {
	var locations map[string]location

	file := getFile("resources/owid-covid-data.json")
	reader := io.Reader(file)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &locations)
	if err != nil {
		log.Fatal(err)
	}

	return locations, []string{}
}

func csvToMaps(reader io.Reader) ([]map[string]string, []string) {
	r := csv.NewReader(reader)
	var rows []map[string]string
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows, header
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

type request struct {
	Countries []string `json:"countries"`
	XStat string `json:"x_stat"`
	YStat string `json:"y_stat"`
}

type line struct {
	Label string
	Data []dataPoint
}

type dataPoint struct {
	Date string
	Value float32
}