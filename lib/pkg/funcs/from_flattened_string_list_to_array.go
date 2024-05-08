package converterfuncs

import (
	"fmt"
	"strings"
)

func FromFlattenedStringListToStringArray(list string, array *[]string) error {
	if array == nil {
		return fmt.Errorf("target array cannot be nil")
	}
	parts := strings.Split(list, ",")
	if len(parts) == 0 {
		*array = []string{}
	} else {
		*array = parts
	}
	return nil
}
