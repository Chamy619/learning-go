package main

import (
	"fmt"
	"strconv"
)

func DoSomeThings(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()

	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}

func doThing1(val int) (int, error) {
	return val, nil
}

func doThing2(val string) (string, error) {
	return val, nil
}

func doThing3(val1 int, val2 string) (string, error) {
	return val2 + strconv.FormatInt(int64(val1), 10), nil
}
