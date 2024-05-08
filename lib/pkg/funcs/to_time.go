package converterfuncs

import (
	"fmt"
	error2 "github.com/smart-libs/go-cc-converter/lib/pkg/error"
	convertertypes "github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"strconv"
	"time"
)

var (
	layoutYyyymmdd     = "20060102"
	layoutYyyymmddSize = len(layoutYyyymmdd)
	layoutYyyyMmDd     = "2006-01-02"
	layoutYyyyMmDdSize = len(layoutYyyyMmDd)
)

func FromStringToDate(str string, dt *time.Time) error {
	switch len(str) {
	case layoutYyyyMmDdSize:
		return fromStringToTimeWithLayout(str, dt, layoutYyyyMmDd)
	case layoutYyyymmddSize:
		return fromStringToTimeWithLayout(str, dt, layoutYyyymmdd)
	default:
		if err := fromStringToTimeWithLayout(str, dt, time.RFC3339Nano); err != nil {
			return err
		}
	}
	*dt = dt.UTC().Truncate(time.Hour)
	return nil
}

func FromStringToTimestamp(str string, dt *time.Time) error {
	return fromStringToTimeWithLayout(str, dt, time.RFC3339Nano)
}

func FromUnixStringToTime[T convertertypes.UnixEpochFormat](from convertertypes.StringUnixTime[T], to *time.Time) error {
	i, err := strconv.ParseInt(string(from), 10, 64)
	if err != nil {
		return err
	}
	var (
		t T
		a any = t
	)
	switch a.(type) {
	case convertertypes.UNIX:
		*to = time.Unix(i, 0)
	case convertertypes.UNIXNano:
		nsec := i % 1000000000
		*to = time.Unix(i, nsec)
	case convertertypes.UNIXMilli:
		*to = time.UnixMilli(i)
	case convertertypes.UNIXMicro:
		*to = time.UnixMicro(i)
	default:
		return error2.NewConversionError(from, to, fmt.Sprintf("unknown format=[%T]", t))
	}
	return nil
}

// fromStringToTimeWithLayout converts a string in the layout given to time.Time
func fromStringToTimeWithLayout(from string, to *time.Time, layout string) error {
	var err error
	if *to, err = time.Parse(layout, from); err != nil {
		return error2.NewConversionErrorWithCause(from, to, err)
	}
	return nil
}

func FromYearStringTimeToInt(from convertertypes.StringTime[convertertypes.RFC3339Year], to *int) error {
	var t time.Time
	err := fromStringToTimeWithLayout(string(from), &t, "2006")
	if err == nil {
		*to = t.Year()
		return nil
	}
	return err
}

func FromStringTimeToTime[T convertertypes.TimeFormat](from convertertypes.StringTime[T], to *time.Time) error {
	var err error
	layouts := convertertypes.GetTimeLayout[T]()
	str := string(from)
	for _, layout := range layouts {
		if err = fromStringToTimeWithLayout(str, to, layout); err == nil {
			return nil
		}
	}
	return err
}
