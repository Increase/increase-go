package core

import (
	"encoding/json"
	"strconv"
)

type LenientInt int

func (li *LenientInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(li))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*li = LenientInt(i)
	return nil
}

type LenientFloat float64

func (li *LenientFloat) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*float64)(li))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*li = LenientFloat(i)
	return nil
}
