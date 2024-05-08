package converterdefault

import (
	funcs "github.com/smart-libs/go-cc-converter/lib/pkg/funcs"
	"github.com/smart-libs/go-cc-converter/lib/pkg/types"
	"strconv"
	"time"
)

func init() {
	AddHandlerWithReturn(strconv.ParseBool)
	AddHandlerWithReturnNoError(strconv.FormatBool)
	AddHandler(funcs.FromAnyToBool)

	AddHandler(func(str convertertypes.FlattenedStringList, array *[]string) error {
		return funcs.FromFlattenedStringListToStringArray(string(str), array)
	})

	AddHandler(funcs.FromStringToDate)
	AddHandler(func(str string, dt *convertertypes.Date) error {
		return funcs.FromStringToDate(str, (*time.Time)(dt))
	})

	AddHandler(funcs.FromYearStringTimeToInt)
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339Year])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339YearMonth])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339Date])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.DateCompressed])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.DateFree])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339WithMin])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339WithFraction])
	AddHandler(funcs.FromStringTimeToTime[convertertypes.RFC3339])

	AddHandler(funcs.FromUnixStringToTime[convertertypes.UNIX])
	AddHandler(funcs.FromUnixStringToTime[convertertypes.UNIXMilli])
	AddHandler(funcs.FromUnixStringToTime[convertertypes.UNIXMicro])
	AddHandler(funcs.FromUnixStringToTime[convertertypes.UNIXNano])
}
