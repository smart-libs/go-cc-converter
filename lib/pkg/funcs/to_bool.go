package converterfuncs

import (
	"fmt"
	"strconv"
)

func FromAnyToBool(from any, to *bool) (err error) {
	*to, err = strconv.ParseBool(fmt.Sprintf("%v", from))
	return err
}
