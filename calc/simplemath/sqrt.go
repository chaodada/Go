package simplemath

import "math"

// 求平方根
func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
