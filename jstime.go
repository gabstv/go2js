package go2js

import (
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
