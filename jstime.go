package go2js

import (
	"database/sql/driver"
	"fmt"
	"reflect"
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
	if val, ok := value.(time.Time); ok {
		*d = NewJsTime(val)
		return nil
	}
	if val, ok := value.(string); ok {
		tt, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		*d = NewJsTime(tt)
		return nil
	}
	if val, ok := value.([]byte); ok {
		tt, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		*d = NewJsTime(tt)
		return nil
	}
	return fmt.Errorf("invalid type %v", reflect.TypeOf(value).String())
}

func (j JsTime) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	tt, _ := time.Parse(RFC3339Milli, string(j))
	return tt.Format("2006-01-02 15:04:05"), nil
}
