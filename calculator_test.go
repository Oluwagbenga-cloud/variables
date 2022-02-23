package calculator

import "testing"

func TestCalculate(t *testing.T) {
	var nums = []struct {
		inp  []string
		outp []float64
	}{
		{[]string{"8*2*5*3", "5+34+7+4", "18-5-3-2", "100/5/2/5"}, []float64{240, 50, 8, 2}},
		{[]string{"3*7*5*9", "50+34+45+4", "150-23-31-2", "1000/5/2/2"}, []float64{945, 133, 94, 50}},
	}

	for _, num := range nums {
		var floatArr = Calculate(num.inp...)
		for i := 0; i < len(num.inp); i++ {
			if floatArr[i] != num.outp[i] {
				t.Errorf("Calculate: Expected %v , found %v", num.outp[i], floatArr[i])
			}
		}
	}
}

func TestAddition(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{3, 4, 5, 3}, 15},
		{[]float64{4, 7, 4, 5}, 20},
	}
	for _, num := range nums {
		result := Addition(num.inp)
		if result != num.outp {
			t.Errorf("Addition: Expected %v found %v", num.outp, result)
		}
	}

}

func TestSubtraction(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{30, 5, 10, 3}, 12},
		{[]float64{40, 7, 4, 5}, 24},
	}
	for _, num := range nums {
		result := Subtraction(num.inp)
		if result != num.outp {
			t.Errorf("Subtraction: Expected %v found %v", num.outp, result)
		}
	}
}

func TestMultiplication(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{3, 4, 5, 3}, 180},
		{[]float64{4, 7, 4, 5}, 560},
	}
	for _, num := range nums {
		result := Multiplication(num.inp)
		if result != num.outp {
			t.Errorf("Multiplication: Expected %v found %v", num.outp, result)
		}
	}

}
func TestDivision(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{400, 4, 5, 2}, 10},
		{[]float64{900, 3, 100, 1}, 3},
	}
	for _, num := range nums {
		result := Division(num.inp)
		if result != num.outp {
			t.Errorf("Division: Expected %v found %v", num.outp, result)
		}
	}
}
