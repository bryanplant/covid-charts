package src

// Common consts
const (
	Countries         = "countries"
	XStat             = "x_stat"
	YStat             = "y_stat"
	CountryFile       = "resources/owid-covid-data.json"
	StateFile         = "resources/states.json"
	StateMetadataFile = "resources/state-metadata.json"
	DateLayout        = "2006-01-02"

	Date = "date"
)

// World data consts
const (
	IsoCode   = "iso_code"
	Continent = "continent"

	TotalCases                 = "total_cases"
	TotalCasesPerMillion       = "total_cases_per_million"
	NewCases                   = "new_cases"
	NewCasesSmoothed           = "new_cases_smoothed"
	NewCasesPerMillion         = "new_cases_per_million"
	NewCasesSmoothedPerMillion = "new_cases_smoothed_per_million"

	TotalDeaths                 = "total_deaths"
	TotalDeathsPerMillion       = "total_deaths_per_million"
	NewDeaths                   = "new_deaths"
	NewDeathsSmoothed           = "new_deaths_smoothed"
	NewDeathsPerMillion         = "new_deaths_per_million"
	NewDeathsSmoothedPerMillion = "new_deaths_smoothed_per_million"

	TotalTests                  = "total_tests"
	TotalTestsPerThousand       = "total_tests_per_thousand"
	NewTests                    = "new_tests"
	NewTestsSmoothed            = "new_tests_smoothed"
	NewTestsPerThousand         = "new_tests_per_thousand"
	NewTestsSmoothedPerThousand = "new_tests_smoothed_per_thousand"

	PositiveRate         = "positive_rate"
	PositiveRateSmoothed = "positive_rate_smoothed"
	TestsPerCase         = "tests_per_case"

	Population        = "population"
	PopulationDensity = "population_density"
	MedianAge         = "median_age"
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
