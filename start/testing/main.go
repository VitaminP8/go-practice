package testing

import (
	"errors"
)

func Division(a, b int) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("division by zero")
	}
	return float64(a) / float64(b), nil
}
