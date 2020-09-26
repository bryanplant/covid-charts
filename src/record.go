package src

import (
	"log"
	"sort"
	"time"
)

const (
	LocationTypeState     = "state"
	LocationTypeCountry   = "country"
	LocationTypeTerritory = "territory"
)

type Location struct {
	Type              string    `json:"type"`
	Continent         *string   `json:"continent"`
	Name              *string   `json:"location"`
	Color             *string   `json:"color"`
	Population        *float64  `json:"population"`
	PopulationDensity *float64  `json:"population_density"`
	MedianAge         *float64  `json:"median_age"`
	Data              []*Record `json:"data"`
}

func (l *Location) populateSmoothedData() {
	sums := []float64{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < len(l.Data); i++ {
		data := l.Data[i]

		// Add next data values
		if data.NewCases != nil {
			sums[0] += *data.NewCases
		}
		if data.NewCasesPerMillion != nil {
			sums[1] += *data.NewCasesPerMillion
		}
		if data.NewDeaths != nil {
			sums[2] += *data.NewDeaths
		}
		if data.NewDeathsPerMillion != nil {
			sums[3] += *data.NewDeathsPerMillion
		}
		if data.NewTests != nil {
			sums[4] += *data.NewTests
		}
		if data.NewTestsPerThousand != nil {
			sums[5] += *data.NewTestsPerThousand
		}
		if data.PositiveRate != nil {
			sums[6] += *data.PositiveRate
		}

		if i >= 6 {
			// Calculate smoothed values
			data.NewCasesSmoothed = getFloat64Pointer(sums[0] / 7)
			data.NewCasesSmoothedPerMillion = getFloat64Pointer(sums[1] / 7)
			data.NewDeathsSmoothed = getFloat64Pointer(sums[2] / 7)
			data.NewDeathsSmoothedPerMillion = getFloat64Pointer(sums[3] / 7)
			data.NewTestsSmoothed = getFloat64Pointer(sums[4] / 7)
			data.NewTestsSmoothedPerThousand = getFloat64Pointer(sums[5] / 7)
			data.PositiveRateSmoothed = getFloat64Pointer(sums[6] / 7)

			data := l.Data[i-6]

			// Subtract old data values
			if data.NewCases != nil {
				sums[0] -= *data.NewCases
			}
			if data.NewCasesPerMillion != nil {
				sums[1] -= *data.NewCasesPerMillion
			}
			if data.NewDeaths != nil {
				sums[2] -= *data.NewDeaths
			}
			if data.NewDeathsPerMillion != nil {
				sums[3] -= *data.NewDeathsPerMillion
			}
			if data.NewTests != nil {
				sums[4] -= *data.NewTests
			}
			if data.NewTestsPerThousand != nil {
				sums[5] -= *data.NewTestsPerThousand
			}
			if data.PositiveRate != nil {
				sums[6] -= *data.PositiveRate
			}
		}
	}
}

type Record struct {
	Date                        jsonDate `json:"date"`
	TotalCases                  *float64 `json:"total_cases"`
	TotalCasesPerMillion        *float64 `json:"total_cases_per_million"`
	NewCases                    *float64 `json:"new_cases"`
	NewCasesSmoothed            *float64 `json:"new_cases_smoothed"`
	NewCasesPerMillion          *float64 `json:"new_cases_per_million"`
	NewCasesSmoothedPerMillion  *float64 `json:"new_cases_smoothed_per_million"`
	TotalDeaths                 *float64 `json:"total_deaths"`
	TotalDeathsPerMillion       *float64 `json:"total_deaths_per_million"`
	NewDeaths                   *float64 `json:"new_deaths"`
	NewDeathsSmoothed           *float64 `json:"new_deaths_smoothed"`
	NewDeathsPerMillion         *float64 `json:"new_deaths_per_million"`
	NewDeathsSmoothedPerMillion *float64 `json:"new_deaths_smoothed_per_million"`
	TotalTests                  *float64 `json:"total_tests"`
	TotalTestsPerThousand       *float64 `json:"total_tests_per_thousand"`
	NewTests                    *float64 `json:"new_tests"`
	NewTestsSmoothed            *float64 `json:"new_tests_smoothed"`
	NewTestsPerThousand         *float64 `json:"new_tests_per_thousand"`
	NewTestsSmoothedPerThousand *float64 `json:"new_tests_smoothed_per_thousand"`
	TestsPerCase                *float64 `json:"tests_per_case"`
	PositiveRate                *float64 `json:"positive_rate"`
	PositiveRateSmoothed        *float64 `json:"positive_rate_smoothed"`
}

func (r Record) getDate() time.Time {
	return time.Time(r.Date)
}

func (r Record) getField(field string) *float64 {
	switch field {
	case TotalCases:
		return r.TotalCases
	case TotalCasesPerMillion:
		return r.TotalCasesPerMillion

	case NewCases:
		return r.NewCases
	case NewCasesSmoothed:
		return r.NewCasesSmoothed
	case NewCasesPerMillion:
		return r.NewCasesPerMillion
	case NewCasesSmoothedPerMillion:
		return r.NewCasesSmoothedPerMillion

	case TotalDeaths:
		return r.TotalDeaths
	case TotalDeathsPerMillion:
		return r.TotalDeathsPerMillion

	case NewDeaths:
		return r.NewDeaths
	case NewDeathsSmoothed:
		return r.NewDeathsSmoothed
	case NewDeathsPerMillion:
		return r.NewDeathsPerMillion
	case NewDeathsSmoothedPerMillion:
		return r.NewDeathsSmoothedPerMillion

	case TotalTests:
		return r.TotalTests
	case TotalTestsPerThousand:
		return r.TotalTestsPerThousand

	case NewTests:
		return r.NewTests
	case NewTestsSmoothed:
		return r.NewTestsSmoothed
	case NewTestsPerThousand:
		return r.NewTestsPerThousand
	case NewTestsSmoothedPerThousand:
		return r.NewTestsSmoothedPerThousand

	case TestsPerCase:
		return r.TestsPerCase
	case PositiveRate:
		return r.PositiveRate
	case PositiveRateSmoothed:
		return r.PositiveRateSmoothed
	}

	log.Fatal("Unknown Record field: " + field)
	return nil
}

type StateRecord struct {
	Date        jsonDate `json:"date"`
	State       *string  `json:"state"`
	TotalCases  *float64 `json:"positive"`
	NewCases    *float64 `json:"positiveIncrease"`
	TotalDeaths *float64 `json:"death"`
	NewDeaths   *float64 `json:"deathIncrease"`
	TotalTests  *float64 `json:"totalTestResults"`
	NewTests    *float64 `json:"totalTestResultsIncrease"`
}

func (r StateRecord) getDate() time.Time {
	return time.Time(r.Date)
}

func StateRecordsToLocations(stateRecords []StateRecord, stateMetadata map[string]StateMetadata) map[string]*Location {
	sort.Slice(stateRecords, func(i, j int) bool {
		return stateRecords[i].getDate().Before(stateRecords[j].getDate())
	})

	locations := map[string]*Location{}
	for _, stateRecord := range stateRecords {
		location, ok := locations[*stateRecord.State]
		if !ok {
			location = &Location{
				Continent: getStringPointer("North America"),
				Name:      stateRecord.State,
				Data:      []*Record{},
			}

			if metadata, ok := stateMetadata[*location.Name]; ok {
				location.Name = metadata.Name
				location.Type = *metadata.Type
				location.Color = metadata.Color
				location.Population = metadata.Population
			}

			locations[*stateRecord.State] = location
		}

		var positiveRate float64
		if stateRecord.NewCases != nil && stateRecord.NewTests != nil && *stateRecord.NewTests != 0 {
			positiveRate = *stateRecord.NewCases / *stateRecord.NewTests
		}

		var totalCasesPerMillion float64
		var newCasesPerMillion float64
		var totalDeathsPerMillion float64
		var newDeathsPerMillion float64
		var totalTestsPerThousand float64
		var newTestsPerThousand float64
		if population := location.Population; population != nil {
			popMillion := *population / 1000000
			popThousand := *population / 1000

			if stateRecord.TotalCases != nil {
				totalCasesPerMillion = *stateRecord.TotalCases / popMillion
			}
			if stateRecord.NewCases != nil {
				newCasesPerMillion = *stateRecord.NewCases / popMillion
			}
			if stateRecord.TotalDeaths != nil {
				totalDeathsPerMillion = *stateRecord.TotalDeaths / popMillion
			}
			if stateRecord.NewDeaths != nil {
				newDeathsPerMillion = *stateRecord.NewDeaths / popMillion
			}
			if stateRecord.TotalTests != nil {
				totalTestsPerThousand = *stateRecord.TotalTests / popThousand
			}
			if stateRecord.NewTests != nil {
				newTestsPerThousand = *stateRecord.NewTests / popThousand
			}
		}

		record := &Record{
			Date:                  stateRecord.Date,
			TotalCases:            stateRecord.TotalCases,
			TotalCasesPerMillion:  &totalCasesPerMillion,
			NewCases:              stateRecord.NewCases,
			NewCasesPerMillion:    &newCasesPerMillion,
			TotalDeaths:           stateRecord.TotalDeaths,
			TotalDeathsPerMillion: &totalDeathsPerMillion,
			NewDeaths:             stateRecord.NewDeaths,
			NewDeathsPerMillion:   &newDeathsPerMillion,
			TotalTests:            stateRecord.TotalTests,
			TotalTestsPerThousand: &totalTestsPerThousand,
			NewTests:              stateRecord.NewTests,
			NewTestsPerThousand:   &newTestsPerThousand,
			PositiveRate:          &positiveRate,
		}
		location.Data = append(location.Data, record)
	}

	return locations
}

type StateMetadata struct {
	Name       *string  `json:"name"`
	Type       *string  `json:"type"`
	Color      *string  `json:"color"`
	Population *float64 `json:"population"`
}

func getStringPointer(s string) *string {
	return &s
}

func getFloat64Pointer(f float64) *float64 {
	return &f
}
