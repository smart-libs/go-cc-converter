package convertertypes

import "time"

/*
 * The date, datetime, time, and duration defined here should be used for parsing and formatting purposes.
 * You should cast string to parse into time.Time and cast time.Time to format it into string.
 */
type (
	/*
	 * ISO-8601 based formats, reference => https://www.w3.org/TR/NOTE-datetime
	 * RFC-3339 => https://datatracker.ietf.org/doc/html/rfc3339
	 */

	// RFC3339Year => Year: YYYY (eg 1997)
	RFC3339Year struct{}
	// RFC3339YearMonth => Year and month: YYYY-MM (eg 1997-07)
	RFC3339YearMonth struct{}
	// RFC3339Date => Complete date: YYYY-MM-DD (eg 1997-07-16)
	RFC3339Date struct{}
	// RFC3339WithMin => Complete date plus hours and minutes: YYYY-MM-DDThh:mmTZD (eg 1997-07-16T19:20+01:00)
	RFC3339WithMin struct{}
	// RFC3339 => Complete date plus hours, minutes and seconds: YYYY-MM-DDThh:mm:ssTZD (eg 1997-07-16T19:20:30+01:00)
	RFC3339 struct{}
	// RFC3339WithFraction => Complete date plus hours, minutes, seconds and a decimal fraction of a second YYYY-MM-DDThh:mm:ss.sTZD (eg 1997-07-16T19:20:30.45+01:00)
	RFC3339WithFraction struct{} //
	// DateCompressed => YYYYMMDD
	DateCompressed struct{}
	// DateFree => DateCompressed or RFC3339Date or YYYY/MM/DD
	DateFree struct{}

	// UNIX is the Unix epoch in seconds
	UNIX struct{}
	// UNIXMilli is the Unix epoch in milliseconds
	UNIXMilli struct{}
	// UNIXMicro is the Unix epoch in microseconds
	UNIXMicro struct{}
	// UNIXNano is the Unix epoch in nanoseconds
	UNIXNano struct{}

	TimeFormat = interface {
		RFC3339Year | RFC3339YearMonth | RFC3339Date | RFC3339WithMin | RFC3339 | RFC3339WithFraction |
			DateCompressed | DateFree
	}

	UnixEpochFormat = interface {
		UNIX | UNIXMilli | UNIXMicro | UNIXNano
	}

	StringTime[F TimeFormat]          string
	StringUnixTime[F UnixEpochFormat] string
)

func GetTimeLayout[F TimeFormat]() []string {
	var (
		f F
		a any = f
	)
	switch a.(type) {
	case RFC3339Year:
		return []string{"2006"}
	case RFC3339YearMonth:
		return []string{"2006-01"}
	case RFC3339Date:
		return []string{"2006-01-02"}
	case RFC3339WithMin:
		return []string{"2006-01-02T15:04", "2006-01-02 15:04"}
	case RFC3339:
		return []string{time.RFC3339}
	case RFC3339WithFraction:
		return []string{time.RFC3339Nano}
	case DateCompressed:
		return []string{"20060102"}
	case DateFree:
		return []string{"2006-01-02", "20060102"}
	default:
		return []string{time.RFC3339Nano}
	}
}
