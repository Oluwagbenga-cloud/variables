package calculator

func Multiplication(arr []float64) float64 {
	var proResult = float64(1)

	for i := 0; i < len(arr); i++ {
		proResult *= arr[i]
	}
	return proResult
}

func Addition(arr []float64) float64 {
	var sumResult float64

	for i := 0; i < len(arr); i++ {
		sumResult += arr[i]
	}
	return sumResult
}

func Subtraction(arr []float64) float64 {
	var diffResult float64 = arr[0]

	for i := 1; i < len(arr); i++ {
		diffResult -= arr[i]
	}
	return diffResult
}

func Division(arr []float64) float64 {
	var divResult float64 = arr[0]

	for i := 1; i < len(arr); i++ {
		divResult /= arr[i]
	}
	return divResult
}
