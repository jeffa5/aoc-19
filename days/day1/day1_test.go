package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1a(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		m int
		e int
	}{
		{
			m: 12, e: 2,
		},
		{
			m: 14, e: 2,
		},
		{
			m: 1969, e: 654,
		},
		{
			m: 100756, e: 33583,
		},
	}

	for _, tt := range tests {
		w := calcWeight(tt.m)
		a.Equal(tt.e, w)
	}
}

func Test1b(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		m int
		e int
	}{
		{
			m: 14, e: 2,
		},
		{
			m: 1969, e: 966,
		},
		{
			m: 100756, e: 50346,
		},
	}

	for _, tt := range tests {
		w := calcWeightWithFuel(tt.m)
		a.Equal(tt.e, w)
	}
}
