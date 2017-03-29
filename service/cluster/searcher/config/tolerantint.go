package config

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// TolerantInt allows the tolerant unmarshalling of both int and string to int
type TolerantInt int

// UnmarshalJSON handles the conversion of string values to int
func (s *TolerantInt) UnmarshalJSON(item []byte) error {
	fmt.Printf("TolerantInt.UnmarshalJSON() is executed with input '%+v'\n", string(item))

	var myString string
	if item[0] == '"' {
		err := json.Unmarshal(item, &myString)
		if err != nil {
			return err
		}
	} else {
		myString = string(item)
	}
	myInt, convErr := strconv.Atoi(myString)
	if convErr != nil {
		return convErr
	}
	*s = (TolerantInt)(myInt)
	return nil
}
