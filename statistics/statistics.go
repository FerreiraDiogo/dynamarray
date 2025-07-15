package statistics

import (
	"errors"
)

const (
	EMPTY_SLICE = "the slice is empty"
)

func Mean(list []int) (float32, error) {
	if len(list) == 0 {
		return 0, errors.New(EMPTY_SLICE)
	}
	acc := 0
	for _, v := range list {
		acc += v
	}
	return float32(acc / len(list)), nil

}

func Median(list []int) (float32, error) {
	if len(list) == 0 {
		return 0, errors.New(EMPTY_SLICE)
	}

	if len(list)%2 == 0 {
		center := list[(len(list) / 2) : (len(list)/2)+1]
		return Mean(center)
	}

	return Mean(list[(len(list) / 2):list[(len(list)/2)]])
}

func Mode(list []int) (int, int, error) {
	if len(list) == 0 {
		return 0, 0, errors.New(EMPTY_SLICE)
	}
	valuesMap := make(map[int]int, 0)
	for _, v := range list {
		valuesMap[v] += 1
	}

	modeKey, modeValue := 0, 0
	for c, v := range valuesMap {
		if valuesMap[c] > modeValue {
			modeValue = v
			modeKey = c
		}
	}

	return modeKey, modeValue, nil
}
