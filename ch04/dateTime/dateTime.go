package dateTime

import "time"

var formats [6]string

const Format string = "2006-01-02 15:04:05"

func init() {
	formats[0] = "2006"
	formats[1] = "2006-01"
	formats[2] = "2006-01-02"
	formats[3] = "2006-01-02 15"
	formats[4] = "2006-01-02 15:04"
	formats[5] = "2006-01-02 15:04:05"
}

func NewDateTime(s string) (time.Time, error) {
	var err error
	var t time.Time

	if s == "" {
		return t, nil
	}

	for _, format := range formats {
		t, err = time.Parse(format, s)
		if err == nil {
			return t, nil
		}
	}
	return t, err
}
