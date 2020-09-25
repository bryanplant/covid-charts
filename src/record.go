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
	Population        *float32  `json:"population"`
	PopulationDensity *float32  `json:"population_density"`
	MedianAge         *float32  `json:"median_age"`
	Data              []*Record `json:"data"`
}

type Record struct {
	Date                        jsonDate `json:"date"`
	TotalCases                  *float32 `json:"total_cases"`
	TotalCasesPerMillion        *float32 `json:"total_cases_per_million"`
	NewCases                    *float32 `json:"new_cases"`
	NewCasesSmoothed            *float32 `json:"new_cases_smoothed"`
	NewCasesPerMillion          *float32 `json:"new_cases_per_million"`
	NewCasesSmoothedPerMillion  *float32 `json:"new_cases_smoothed_per_million"`
	TotalDeaths                 *float32 `json:"total_deaths"`
	TotalDeathsPerMillion       *float32 `json:"total_deaths_per_million"`
	NewDeaths                   *float32 `json:"new_deaths"`
	NewDeathsSmoothed           *float32 `json:"new_deaths_smoothed"`
	NewDeathsPerMillion         *float32 `json:"new_deaths_per_million"`
	NewDeathsSmoothedPerMillion *float32 `json:"new_deaths_smoothed_per_million"`
	TotalTests                  *float32 `json:"total_tests"`
	TotalTestsPerThousand       *float32 `json:"total_tests_per_thousand"`
	NewTests                    *float32 `json:"new_tests"`
	NewTestsSmoothed            *float32 `json:"new_tests_smoothed"`
	NewTestsPerThousand         *float32 `json:"new_tests_per_thousand"`
	NewTestsSmoothedPerThousand *float32 `json:"new_tests_smoothed_per_thousand"`
	TestsPerCase                *float32 `json:"tests_per_case"`
	PositiveRate                *float32 `json:"positive_rate"`
	PositiveRateSmoothed        *float32 `json:"positive_rate_smoothed"`
}

func (r Record) getDate() time.Time {
	return time.Time(r.Date)
}

func (r Record) getField(field string) *float32 {
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
	TotalCases  *float32 `json:"positive"`
	NewCases    *float32 `json:"positiveIncrease"`
	TotalDeaths *float32 `json:"death"`
	NewDeaths   *float32 `json:"deathIncrease"`
	TotalTests  *float32 `json:"totalTestResults"`
	NewTests    *float32 `json:"totalTestResultsIncrease"`
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

		var positiveRate float32
		if stateRecord.NewCases != nil && stateRecord.NewTests != nil && *stateRecord.NewTests != 0 {
			positiveRate = *stateRecord.NewCases / *stateRecord.NewTests
		}

		var totalCasesPerMillion float32
		var newCasesPerMillion float32
		var totalDeathsPerMillion float32
		var newDeathsPerMillion float32
		var totalTestsPerThousand float32
		var newTestsPerThousand float32
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

		var positiveRateSmoothed float32
		if len(location.Data) >= 7 {
			var sum float32
			for i := 1; i <= 7; i++ {
				value := location.Data[len(location.Data)-i].PositiveRate
				if value != nil {
					sum += *value
				}
			}
			positiveRateSmoothed = sum / 7
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
			PositiveRateSmoothed:  &positiveRateSmoothed,
		}
		location.Data = append(location.Data, record)
	}

	return locations
}

type StateMetadata struct {
	Name       *string  `json:"name"`
	Type       *string  `json:"type"`
	Color      *string  `json:"color"`
	Population *float32 `json:"population"`
}

func getStringPointer(s string) *string {
	return &s
}
