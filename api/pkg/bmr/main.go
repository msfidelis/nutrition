package bmr

import "strings"

func Calc(gender string, weight float64, height float64, age int) float64 {
	var basal float64
	if strings.ToUpper(gender) == "M" {
		basal = 66 + (13.75 * weight) + (5.0 * (height * 100)) - (6.8 * float64(age))
	} else {
		basal = 665 + (9.56 * weight) + (1.8 * (height * 100)) - (4.7 * float64(age))
	}
	return basal
}
