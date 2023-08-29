package calcy

import "errors"

type Calculator interface {
	Calculate(int, int) (int, error)
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}
type Bogus struct{ Offset int }

func (this Addition) Calculate(a, b int) (int, error)       { return a + b, nil }
func (this Subtraction) Calculate(a, b int) (int, error)    { return a - b, nil }
func (this Multiplication) Calculate(a, b int) (int, error) { return a * b, nil }
func (this Division) Calculate(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
func (this Bogus) Calculate(a, b int) (int, error) { return this.Offset + a + b, nil }

var ErrDivideByZero = errors.New("can't divide by zero")
