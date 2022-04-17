package helpers

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Datetime struct {
	time.Time
}

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
	// first we get the bytes from the original value
	bytes, ok := value.([]byte)
  if !ok {
    return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
  }
	// here we will parse the data into the custom data type we want
	strInput := strings.Trim(string(bytes), `"`)
	if newTime, err := time.Parse("2006-01-02", strInput); err != nil {
		// then we set the value of the custom data type to the struct
		j.Time = newTime
	}
	return nil
}
// here we just return the value
func (j Datetime) Value() (driver.Value, error) {
	return j.Time, nil
}
