package calcy

import "testing"

func TestAdditionSuccess(t *testing.T) {
	result := Addition{}.Calculate(5, 4)
	if result != 9 {
		t.Error("want 9, got", result)
	}
}

func TestAdditionNegativeParameter(t *testing.T) {
	result := Addition{}.Calculate(5, -4)
	if result != 1 {
		t.Error("want 1, got", result)
	}
}

func TestAdditionNegativeResult(t *testing.T) {
	result := Addition{}.Calculate(-2, -4)
	if result != -6 {
		t.Error("want -6, got", result)
	}
}

func TestAdditionLargeNumbers(t *testing.T) {
	result := Addition{}.Calculate(1000000000, 1000000000)
	if result != 2000000000 {
		t.Error("want 2000000000, got", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction{}.Calculate(5, 4)
	if result != 1 {
		t.Error("want 1, got", result)
	}
}

func TestSubtractionNegativeParameter(t *testing.T) {
	result := Subtraction{}.Calculate(5, -4)
	if result != 9 {
		t.Error("want 9, got", result)
	}
}

func TestSubtractionNegativeResult(t *testing.T) {
	result := Subtraction{}.Calculate(-2, 4)
	if result != -6 {
		t.Error("want -6, got", result)
	}
}

func TestSubtractionLargeNumbers(t *testing.T) {
	result := Subtraction{}.Calculate(2000000000, 1000000000)
	if result != 1000000000 {
		t.Error("want 1000000000, got", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication{}.Calculate(4, 5)
	if result != 20 {
		t.Error("want 20, got", result)
	}
}

func TestMultiplicationNegativeParameter(t *testing.T) {
	result := Multiplication{}.Calculate(-5, -4)
	if result != 20 {
		t.Error("want 20, got", result)
	}
}

func TestMultiplicationNegativeResult(t *testing.T) {
	result := Multiplication{}.Calculate(-2, 4)
	if result != -8 {
		t.Error("want -8, got", result)
	}
}

func TestMultiplicationLargeNumbers(t *testing.T) {
	result := Multiplication{}.Calculate(100000, 100000)
	if result != 10000000000 {
		t.Error("want 10000000000, got", result)
	}
}

func TestDivision(t *testing.T) {
	result := Division{}.Calculate(4, 5)
	if result != 0 {
		t.Error("want 0, got", result)
	}
}

func TestDivisionNegativeParameter(t *testing.T) {
	result := Division{}.Calculate(-10, -2)
	if result != 5 {
		t.Error("want 5, got", result)
	}
}

func TestDivisionNegativeResult(t *testing.T) {
	result := Division{}.Calculate(4, -2)
	if result != -2 {
		t.Error("want -2, got", result)
	}
}

func TestDivisionLargeNumbers(t *testing.T) {
	result := Division{}.Calculate(1000000000, 1000000000)
	if result != 1 {
		t.Error("want 1, got", result)
	}
}

func TestBogus(t *testing.T) {
	result := Bogus{Offset: 42}.Calculate(4, 5)
	if result != 42+9 {
		t.Error("want 51, got", result)
	}
}

func TestBogusNegativeParameter(t *testing.T) {
	result := Bogus{Offset: -42}.Calculate(65, -4)
	if result != 19 {
		t.Error("want 19, got", result)
	}
}

func TestBogusNegativeResult(t *testing.T) {
	result := Bogus{Offset: -2}.Calculate(-2, -4)
	if result != -8 {
		t.Error("want -8, got", result)
	}
}

func TestBogusLargeNumbers(t *testing.T) {
	result := Bogus{Offset: 1000000000}.Calculate(1000000000, 1000000000)
	if result != 3000000000 {
		t.Error("want 2000000000, got", result)
	}
}
