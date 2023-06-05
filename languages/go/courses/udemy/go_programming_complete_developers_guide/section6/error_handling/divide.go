package main

import "errors"

func divide(lhs int, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		return lhs / rhs, nil
	}
}
