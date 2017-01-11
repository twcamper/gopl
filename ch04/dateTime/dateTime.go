package dateTime

import "time"

type DateTime struct {
	format string
	value  time.Time
}

var formats [6]string

func init() {
	formats[0] = "2006"
	formats[1] = "2006-01"
	formats[2] = "2006-01-02"
	formats[3] = "2006-01-02 15"
	formats[4] = "2006-01-02 15:04"
	formats[5] = "2006-01-02 15:04:05"
}

func NewDateTime(s string) (*DateTime, error) {
	var err error
	var t time.Time
	p := new(DateTime)

	for _, format := range formats {
		t, err = time.Parse(format, s)
		if err == nil {
			p.format = format
			p.value = t
			return p, nil
		}
	}
	return p, err
}

func (dt *DateTime) String() string {
	return dt.value.Format(dt.format)
}

func (dt *DateTime) Format() string {
	return dt.format
}

func (dt *DateTime) Value() time.Time {
	return dt.value
}
