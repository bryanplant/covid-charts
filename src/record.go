package src

import (
	"log"
)

type location struct {
	Continent         *string   `json:"continent"`
	Location          *string   `json:"location"`
	Population        *float32  `json:"population"`
	PopulationDensity *float32  `json:"population_density"`
	MedianAge         *float32  `json:"median_age"`
	Data              []*record `json:"data"`
}

type record struct {
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
	TestsPerCase                *float32 `json:"positive_rate"`
	PositiveRate                *float32 `json:"tests_per_case"`
}

func (r record) getField(field string) *float32 {
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
	}

	log.Fatal("Unknown record field: " + field)
	return nil
}

//func parseMetadata(m map[string]string) location {
//	return location{
//		Continent:         getStringPointer(m[Continent]),
//		Location:          getStringPointer(m[Location]),
//		Population:        parseFloat(m[Population]),
//		PopulationDensity: parseFloat(m[PopulationDensity]),
//	}
//}
//
//func parseRecord(m map[string]string) record {
//	return record{
//		Date:                        parseDate(m[Date]),
//		TotalCases:                  parseFloat(m[TotalCases]),
//		TotalCasesPerMillion:        parseFloat(m[TotalCasesPerMillion]),
//		NewCases:                    parseFloat(m[NewCases]),
//		NewCasesSmoothed:            parseFloat(m[NewCasesSmoothed]),
//		NewCasesPerMillion:          parseFloat(m[NewCasesPerMillion]),
//		NewCasesSmoothedPerMillion:  parseFloat(m[NewCasesSmoothedPerMillion]),
//		TotalDeaths:                 parseFloat(m[TotalDeaths]),
//		TotalDeathsPerMillion:       parseFloat(m[TotalDeathsPerMillion]),
//		NewDeaths:                   parseFloat(m[NewDeaths]),
//		NewDeathsSmoothed:           parseFloat(m[NewDeathsSmoothed]),
//		NewDeathsPerMillion:         parseFloat(m[NewDeathsPerMillion]),
//		NewDeathsSmoothedPerMillion: parseFloat(m[NewDeathsSmoothedPerMillion]),
//		TotalTests:                  parseFloat(m[TotalTests]),
//		TotalTestsPerThousand:       parseFloat(m[TotalTestsPerThousand]),
//		NewTests:                    parseFloat(m[NewTests]),
//		NewTestsSmoothed:            parseFloat(m[NewTestsSmoothed]),
//		NewTestsPerThousand:         parseFloat(m[NewTestsPerThousand]),
//		NewTestsSmoothedPerThousand: parseFloat(m[NewTestsSmoothedPerThousand]),
//		TestsPerCase:                parseFloat(m[TestsPerCase]),
//		PositiveRate:                parseFloat(m[PositiveRate]),
//	}
//}
//
//func parseDate(dateString string) time.Time {
//	date, err := time.Parse(DateLayout, dateString)
//	if err != nil {
//		log.Fatalln("Failed to parse date: "+dateString, err)
//	}
//
//	return date
//}
//
//func parseFloat(floatString string) *float32 {
//	if floatString == "" {
//		return nil
//	}
//
//	f, err := strconv.ParseFloat(floatString, 32)
//	if err != nil {
//		log.Fatalln("Failed to parse float: "+floatString, err)
//	}
//
//	f32 := float32(f)
//	return &f32
//}
//
//func getStringPointer(s string) *string {
//	return &s
//}
