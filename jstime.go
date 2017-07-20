package go2js

import (
	"database/sql/driver"
	"time"
)

const (
	JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"
	RFC3339Milli        = "2006-01-02T15:04:05.999Z07:00"
)

func NewJsTime(t time.Time) JsTime {
	return JsTime(t.Format(RFC3339Milli))
}

type JsTime string

func (d *JsTime) Scan(value interface{}) error {
	if value == nil {
		*d = ""
		return nil
	}
	vv, ok := value.(string)
	if ok {
		tt, err := time.Parse("2006-01-02 15:04:05", vv)
		if err != nil {
			return err
		}
		*d = NewJsTime(tt)
	} else {
		vv2, _ := value.([]byte)
		tt, err := time.Parse("2006-01-02 15:04:05", string(vv2))
		if err != nil {
			return err
		}
		*d = NewJsTime(tt)
	}
	return nil
}

func (j JsTime) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	tt, _ := time.Parse(RFC3339Milli, string(j))
	return tt.Format("2006-01-02 15:04:05"), nil
}
