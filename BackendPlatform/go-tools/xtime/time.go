package xtime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(xt.Duration(value))
		return nil
	case string:
		t, err := xt.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(t)
		return nil
	default:
		return errors.New("invalid duration")
	}
}
