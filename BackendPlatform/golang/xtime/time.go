package xtime

import (
	"database/sql/driver"
	"strconv"
	xt "time"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (t *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xt.Time:
		*t = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*t = Time(i)
	}
	return
}

// Value get time value.
func (t Time) Value() (driver.Value, error) {
	return xt.Unix(int64(t), 0), nil
}

// Time get time.
func (t Time) Time() xt.Time {
	return xt.Unix(int64(t), 0)
}

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration xt.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xt.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}
