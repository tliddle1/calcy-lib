package calcy

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"10+5": {
			a:        10,
			b:        5,
			expected: 15,
		},
		"(-2)+(-4)": {
			a:        -2,
			b:        -4,
			expected: -6,
		},
		"5+(-4)": {
			a:        5,
			b:        -4,
			expected: 1,
		},
		"1000000000+1000000000": {
			a:        1000000000,
			b:        1000000000,
			expected: 2000000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Addition{}.Calculate(test.a, test.b)
			if err != nil {
				t.Error(err)
			}
			if actual != test.expected {
				t.Errorf("sums don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"10-6": {
			a:        10,
			b:        6,
			expected: 4,
		},
		"(-2)-(-4)": {
			a:        -2,
			b:        -4,
			expected: 2,
		},
		"5-(-4)": {
			a:        5,
			b:        -4,
			expected: 9,
		},
		"1000000000-2000000000": {
			a:        1000000000,
			b:        2000000000,
			expected: -1000000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Subtraction{}.Calculate(test.a, test.b)
			if err != nil {
				t.Error(err)
			}
			if actual != test.expected {
				t.Errorf("results don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"5*4": {
			a:        5,
			b:        4,
			expected: 20,
		},
		"(-5)*(-4)": {
			a:        -5,
			b:        -4,
			expected: 20,
		},
		"(-2)*4": {
			a:        -2,
			b:        4,
			expected: -8,
		},
		"100000*100000": {
			a:        100000,
			b:        100000,
			expected: 10000000000,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Multiplication{}.Calculate(test.a, test.b)
			if err != nil {
				t.Error(err)
			}
			if actual != test.expected {
				t.Errorf("products don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"5/4": {
			a:        5,
			b:        4,
			expected: 1,
		},
		"(-10)/(-2)": {
			a:        -10,
			b:        -2,
			expected: 5,
		},
		"4/(-2)": {
			a:        4,
			b:        -2,
			expected: -2,
		},
		"1000000000/1000000000": {
			a:        1000000000,
			b:        1000000000,
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Division{}.Calculate(test.a, test.b)
			if err != nil {
				t.Error(err)
			}
			if actual != test.expected {
				t.Errorf("quotients don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestDivisionByZero(t *testing.T) {
	t.Parallel()
	actual, err := Division{}.Calculate(34, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("expected divide by zero error but got %v", err)
	}
	fmt.Println(actual)
}

func TestBogus(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		a        int
		b        int
		offset   int
		expected int
	}{
		"5/4": {
			a:        5,
			b:        4,
			offset:   42,
			expected: 51,
		},
		"(-10)/(-2)": {
			a:        -10,
			b:        -2,
			offset:   -31,
			expected: -43,
		},
		"4/(-2)": {
			a:        4,
			b:        -2,
			offset:   52,
			expected: 54,
		},
		"1000000000/1000000000": {
			a:        1000000000,
			b:        1000000000,
			offset:   1,
			expected: 2000000001,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Bogus{Offset: test.offset}.Calculate(test.a, test.b)
			if err != nil {
				t.Error(err)
			}
			if actual != test.expected {
				t.Errorf("results don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
