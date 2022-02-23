package calculator

import (
	"strconv"
	"strings"
)

func Calculate(g ...string) []float64 {
	var addition []float64
	var multiply []float64
	var divide []float64
	var subtract []float64
	var totalResult []float64

	for i := 0; i < len(g); i++ {

		if strings.Contains(g[i], "+") == true {
			additionCon := strings.Split(g[i], "+")

			for i := 0; i < len(additionCon); i++ {
				a, _ := strconv.Atoi(additionCon[i])
				addition = append(addition, float64(a))
			}
			findAddition := Addition(addition)
			totalResult = append(totalResult, findAddition)

		} else if strings.Contains(g[i], "*") == true {
			multiplyCon := strings.Split(g[i], "*")

			for i := 0; i < len(multiplyCon); i++ {
				m, _ := strconv.Atoi(multiplyCon[i])
				multiply = append(multiply, float64(m))

			}
			findMultiplication := Multiplication(multiply)
			totalResult = append(totalResult, findMultiplication)

		} else if strings.Contains(g[i], "/") == true {
			divisionCon := strings.Split(g[i], "/")

			for i := 0; i < len(divisionCon); i++ {
				d, _ := strconv.Atoi(divisionCon[i])
				divide = append(divide, float64(d))
			}
			findDivision := Division(divide)
			totalResult = append(totalResult, findDivision)

		} else if strings.Contains(g[i], "-") == true {
			subtractCon := strings.Split(g[i], "-")

			for i := 0; i < len(subtractCon); i++ {
				s, _ := strconv.Atoi(subtractCon[i])
				subtract = append(subtract, float64(s))
			}
			findSubtraction := Subtraction(subtract)
			totalResult = append(totalResult, findSubtraction)
		}
	}
	return totalResult
}
