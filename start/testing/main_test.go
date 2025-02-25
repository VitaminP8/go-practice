package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	testCases := []struct {
		num  int
		div  int
		want float64
	}{
		{10, 1, 10.0},
		{20, 10, 2.0},
		{25, 25, 1.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(" %d/%d", tc.num, tc.div), func(t *testing.T) {
			got, err := Division(tc.num, tc.div)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := Division(10, 0)
	assert.Error(t, err)
}
