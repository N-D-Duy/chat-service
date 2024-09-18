package utils

import "strconv"

func StringToUint(s string) (uint, error) {
	uintValue, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uintValue), nil

}
