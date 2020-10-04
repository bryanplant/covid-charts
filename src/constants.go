package src

// Common consts
const (
	XStat        = "x_stat"
	YStat        = "y_stat"
	CountryStats = "country_stats"
	StateStats   = "state_stats"

	CountryURL = "https://covid.ourworldindata.org/data/owid-covid-data.json"
	StateURL   = "https://api.covidtracking.com/v1/states/daily.json"
	DateLayout = "2006-01-02"
)

// World data consts
const (
	IsoCode   = "iso_code"
	Continent = "continent"

	TotalCases                 = "Total Cases"
	TotalCasesPerMillion       = "Total Cases Per Million"
	NewCases                   = "New Cases"
	NewCasesSmoothed           = "New Cases Smoothed"
	NewCasesPerMillion         = "New Cases Per Million"
	NewCasesSmoothedPerMillion = "New Cases Smoothed Per Million"

	TotalDeaths                 = "Total Deaths"
	TotalDeathsPerMillion       = "Total Deaths Per Million"
	NewDeaths                   = "New Deaths"
	NewDeathsSmoothed           = "New Deaths Smoothed"
	NewDeathsPerMillion         = "New Deaths Per Million"
	NewDeathsSmoothedPerMillion = "New Deaths Smoothed Per Million"

	TotalTests                  = "Total Tests"
	TotalTestsPerThousand       = "Total Tests Per Thousand"
	NewTests                    = "New Tests"
	NewTestsSmoothed            = "New Tests Smoothed"
	NewTestsPerThousand         = "New Tests Per Thousand"
	NewTestsSmoothedPerThousand = "New Tests Smoothed Per Thousand"

	PositiveRate         = "Positive Rate"
	PositiveRateSmoothed = "Positive Rate Smoothed"
	TestsPerCase         = "Tests Per Case"

	Population        = "Population"
	PopulationDensity = "Population Density"
	MedianAge         = "Median Age"
)

// US data consts
const (
	State                            = "state"
	DataQualityGrade                 = "dataQualityGrade"
	Death                            = "death"
	DeathConfirmed                   = "deathConfirmed"
	DeathIncrease                    = "deathIncrease"
	DeathProbable                    = "deathProbable"
	Hospitalized                     = "hospitalized"
	HospitalizedCumulative           = "hospitalizedCumulative"
	HospitalizedCurrently            = "hospitalizedCurrently"
	HospitalizedIncrease             = "hospitalizedIncrease"
	InIcuCumulative                  = "inIcuCumulative"
	InIcuCurrently                   = "inIcuCurrently"
	Negative                         = "Negative"
	NegativeIncrease                 = "negativeIncrease"
	NegativeTestsAntibody            = "negativeTestsAntibody"
	NegativeTestsPeopleAntibody      = "negativeTestsPeopleAntibody"
	NegativeTestsViral               = "negativeTestsViral"
	OnVentilatorCumulative           = "onVentilatorCumulative"
	OnVentilatorCurrently            = "onVentilatorCurrently"
	Pending                          = "pending"
	Positive                         = "positive"
	PositiveCasesViral               = "positiveCasesViral"
	PositiveIncrease                 = "positiveIncrease"
	PositiveScore                    = "positiveScore"
	PositiveTestsAntibody            = "positiveTestsAntibody"
	PositiveTestsAntigen             = "positiveTestsAntigen"
	PositiveTestsPeopleAntibody      = "positiveTestsPeopleAntibody"
	PositiveTestsPeopleAntigen       = "positiveTestsPeopleAntigen"
	PositiveTestsViral               = "positiveTestsViral"
	Recovered                        = "recovered"
	TotalTestEncountersViral         = "totalTestEncountersViral"
	TotalTestEncountersViralIncrease = "totalTestEncountersViralIncrease"
	TotalTestResults                 = "totalTestResults"
	TotalTestResultsIncrease         = "totalTestResultsIncrease"
	TotalTestsAntibody               = "totalTestsAntibody"
	TotalTestsAntigen                = "totalTestsAntigen"
	TotalTestsPeopleAntibody         = "totalTestsPeopleAntibody"
	TotalTestsPeopleAntigen          = "totalTestsPeopleAntigen"
	TotalTestsPeopleViral            = "totalTestsPeopleViral"
	TotalTestsPeopleViralIncrease    = "totalTestsPeopleViralIncrease"
	TotalTestsViral                  = "totalTestsViral"
	TotalTestsViralIncrease          = "totalTestsViralIncrease"
)
