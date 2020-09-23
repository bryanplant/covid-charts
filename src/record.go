package src

import (
	"log"
	"strconv"
	"time"
)

type record struct {
	IsoCode                     string
	Continent                   string
	Location                    string
	Date                        time.Time
	TotalCases                  *float32
	TotalCasesPerMillion        *float32
	NewCases                    *float32
	NewCasesSmoothed            *float32
	NewCasesPerMillion          *float32
	NewCasesSmoothedPerMillion  *float32
	TotalDeaths                 *float32
	TotalDeathsPerMillion       *float32
	NewDeaths                   *float32
	NewDeathsSmoothed           *float32
	NewDeathsPerMillion         *float32
	NewDeathsSmoothedPerMillion *float32
	TotalTests                  *float32
	TotalTestsPerThousand       *float32
	NewTests                    *float32
	NewTestsSmoothed            *float32
	NewTestsPerThousand         *float32
	NewTestsSmoothedPerThousand *float32
	TestsPerCase                *float32
	PositiveRate                *float32
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

func parseRecord(m map[string]string) record {
	return record{
		IsoCode:   m[IsoCode],
		Continent: m[Continent],
		Location:  m[Location],
		Date:      parseDate(m[Date]),

		TotalCases:           parseFloat(m[TotalCases]),
		TotalCasesPerMillion: parseFloat(m[TotalCasesPerMillion]),

		NewCases:                   parseFloat(m[NewCases]),
		NewCasesSmoothed:           parseFloat(m[NewCasesSmoothed]),
		NewCasesPerMillion:         parseFloat(m[NewCasesPerMillion]),
		NewCasesSmoothedPerMillion: parseFloat(m[NewCasesSmoothedPerMillion]),

		TotalDeaths:           parseFloat(m[TotalDeaths]),
		TotalDeathsPerMillion: parseFloat(m[TotalDeathsPerMillion]),

		NewDeaths:                   parseFloat(m[NewDeaths]),
		NewDeathsSmoothed:           parseFloat(m[NewDeathsSmoothed]),
		NewDeathsPerMillion:         parseFloat(m[NewDeathsPerMillion]),
		NewDeathsSmoothedPerMillion: parseFloat(m[NewDeathsSmoothedPerMillion]),

		TotalTests:            parseFloat(m[TotalTests]),
		TotalTestsPerThousand: parseFloat(m[TotalTestsPerThousand]),

		NewTests:                    parseFloat(m[NewTests]),
		NewTestsSmoothed:            parseFloat(m[NewTestsSmoothed]),
		NewTestsPerThousand:         parseFloat(m[NewTestsPerThousand]),
		NewTestsSmoothedPerThousand: parseFloat(m[NewTestsSmoothedPerThousand]),

		TestsPerCase: parseFloat(m[TestsPerCase]),
		PositiveRate: parseFloat(m[PositiveRate]),
	}
}

func parseDate(dateString string) time.Time {
	date, err := time.Parse(DateLayout, dateString)
	if err != nil {
		log.Fatalln("Failed to parse date: "+dateString, err)
	}

	return date
}

func parseFloat(floatString string) *float32 {
	if floatString == "" {
		return nil
	}

	f, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		log.Fatalln("Failed to parse float: "+floatString, err)
	}

	f32 := float32(f)
	return &f32
}
