package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	if res, err := MySqrt(-4); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	if res, err := MySqrtAnonymous(4); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func MySqrt(val float64) (result float64, err error) {
	if val < 0 {
		result = 0
		err = errors.New("Can not calculate a negative number")
	} else {
		result = math.Sqrt(val)
		err = nil
	}

	return
}

func MySqrtAnonymous(val float64) (float64, error) {
	if val < 0 {
		return math.Sqrt(val), errors.New("Can not calculate a negative number")
	} else {
		return math.Sqrt(val), nil
	}
}
