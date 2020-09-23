package src

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"

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

	var countries []string
	for country := range data {
		countries = append(countries, country)
	}

	sort.Strings(countries)
	sort.Strings(stats)

	options := map[string][]string{}
	options["countries"] = countries
	options["stats"] = stats

	bytes, err := json.Marshal(countries)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func getLine(data map[string][]record, country, stat string) line {
	countryData := data[country]

	var dataPoints []dataPoint
	for _, item := range countryData {
		if item.getField(stat) != nil {
			dataPoint := dataPoint{
				Date:  item.Date.Format(DateLayout),
				Value: *item.getField(stat),
			}
			dataPoints = append(dataPoints, dataPoint)
		}
	}

	return line{
		Label: country,
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

func readData() (map[string][]record, []string) {
	countries := map[string][]record{}

	csvFile := getFile(DataFile)
	reader := io.Reader(csvFile)

	maps, stats := csvToMaps(reader)
	for _, m := range maps {
		record := parseRecord(m)
		countries[record.Location] = append(countries[record.Location], record)
	}

	return countries, stats
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
	var csvFile *os.File
	var err error

	csvFile, err = os.Open(path)
	if err != nil {
		csvFile, err = os.Open("src/" + path)
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
	}

	return csvFile
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