package src

import (
	"strings"
	"time"
)

type jsonDate time.Time

func (j *jsonDate) UnmarshalJSON(b []byte) error {
	var t time.Time
	var err error

	s := strings.Trim(string(b), "\"")
	t, err = time.Parse("2006-01-02", s)
	if err != nil {
		t, err = time.Parse("20060102", s)
		if err != nil {
			return err
		}
	}
	*j = jsonDate(t)
	return nil
}

func (j *jsonDate) String() string {
	return time.Time(*j).Format(DateLayout)
}
