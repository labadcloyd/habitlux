package helpers

import (
	"database/sql/driver"
	"strings"
	"time"
)

type Datetime struct {
	time.Time
}

// for parsing datetime on fiber http request
func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}

// for gorm custom data type
// the Scan method will try to convert the raw data to the data type gorm accepts
func (j *Datetime) Scan(value interface{}) error {
	if t, ok := value.(time.Time); ok {
		j.Time = t
	}
	return nil
}

// here we just return the value
func (j Datetime) Value() (driver.Value, error) {
	return j.Time, nil
}
