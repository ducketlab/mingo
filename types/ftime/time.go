package ftime

import (
	"strconv"
	"time"
)

type FormatType string

const (
	TIMESTAMP = FormatType("timestamp")
	TEXT = FormatType("text")
)

type TimestampLength int

const (
	Length10 = TimestampLength(10)
	Length13 = TimestampLength(13)
	Length16 = TimestampLength(16)
	Length19 = TimestampLength(19)
)

var (
	UsedFormatType = TIMESTAMP
	UsedTimestampLength = Length13
	UsedTextTimeFormat = "2006-01-02 15:04:05"
)

func Now() Time {
	return T(time.Now())
}

func T(t time.Time) Time {
	return Time(t)
}

type Time time.Time

func (t *Time) T() time.Time {
	return time.Time(*t)
}

func (t *Time) parseTS(ts string) error {

	for len(ts) < 19 {
		ts += "0"
	}

	sec, err := strconv.ParseInt(ts[:10], 10, 64)
	if err != nil {
		return err
	}

	nsec, err := strconv.ParseInt(ts[10:], 10, 64)

	*t = Time(time.Unix(sec, nsec))
	return nil
}

func (t *Time) parseTSInt64(ts int64) error {
	return t.parseTS(strconv.FormatInt(ts, 10))
}

func (t *Time) parseText(data string) error {
	now, err := time.ParseInLocation(`"`+UsedTextTimeFormat+`"`, data, time.Local)
	if err != nil {
		return err
	}
	*t = Time(now)
	return nil
}

func (t Time) formatText() []byte {
	b := make([]byte, 0, len(UsedTextTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, UsedTextTimeFormat)
	b = append(b, '"')
	return b
}

func (t Time) Timestamp() int64 {
	var ts int64

	if t.T().IsZero() {
		return 0
	}

	switch UsedTimestampLength {
	case Length10:
		ts = time.Time(t).Unix()
	case Length13:
		ts = time.Time(t).UnixNano() / 1000000
	case Length16:
		ts = time.Time(t).UnixNano() / 1000
	case Length19:
		ts = time.Time(t).UnixNano()
	default:
		ts = time.Time(t).UnixNano() / 1000000
	}

	return ts
}
