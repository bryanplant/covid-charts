package src

import (
	"strings"
	"time"
)

type jsonDate time.Time

func (j *jsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = jsonDate(t)
	return nil
}
