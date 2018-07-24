package sample

import "errors"

func Add(x int, y int) (z int, err error) {
	return x + y, nil
}

func Subtract(x int, y int) (z int, err error) {
	return  x - y, nil
}

func Multiply(x int, y int) (z int, err error) {
	return x * y, nil
}

func Divide(x int, y int) (z int, err error) {
	if y == 0 {
		return 0, errors.New("y != 0")
	} else {
		return x / y, nil
	}
}