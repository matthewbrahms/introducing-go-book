package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMath(t *testing.T) {
	cases := []struct {
		xs       []float64
		max, min float64
	}{
		{
			xs:  []float64{3, 5, 2, 1, 7, 9},
			max: 9,
			min: 1,
		},
		{
			xs:  []float64{},
			max: 0,
			min: 0,
		},
	}

	for _, c := range cases {
		max := Max(c.xs)
		if max != c.max {
			t.Errorf("Expected %f got %f", c.max, max)
		}
		min := Min(c.xs)
		if min != c.min {
			t.Errorf("Expected %f got %f", c.min, min)
		}
	}
}
