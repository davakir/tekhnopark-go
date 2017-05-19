package main

import (
	"math"
)

func Sqrt(x float64) float64 {
	number := int64(NewtonSqrt(x, x))
	return float64(number) / 100.0
}


// Вычисление квадратного корня числа Методом Ньютона
func NewtonSqrt(a float64, x0 float64) float64 {
	var e float64 = 0.01
	var xk float64 = x0
	var xk1 float64

	if a == 0 {
		return 0
	} else if a < 0 {
		return math.NaN()
	}

	for {
		xk1 = (1.0 / 2.0) * (xk + a / xk)

		if xk1 < 0 {
			xk1 = -xk1
		}

		var tmp float64 = xk1 - xk
		if tmp < 0 {
			tmp = -tmp
		}

		if tmp < e {
			return xk1 * 100
		} else {
			xk = xk1
		}
	}


}
