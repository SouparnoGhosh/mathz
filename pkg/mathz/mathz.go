package mathz

import "strconv"

func Add(v1, v2 string) (result string) {
	i1, i2 := errHandler(v1, v2)

	result = strconv.Itoa(i1 + i2)
	return result
}

func Multiply(v1, v2 string) (result string) {
	i1, i2 := errHandler(v1, v2)

	result = strconv.Itoa(i1 * i2)
	return result
}

func Subtract(v1, v2 string) (result string) {
	i1, i2 := errHandler(v1, v2)

	result = strconv.Itoa(i1 - i2)
	return result
}

func Divide(v1, v2 string) (result string) {
	i1, i2 := errHandler(v1, v2)
	if i2 == 0 {
		result = "Can't divide by 0"
		return result
	}

	result = strconv.FormatFloat(float64(i1) / float64(i2), 'f', 5, 32)
	return result
}

// Helper function

func errHandler(v1, v2 string) (i1, i2 int) {
	i1, err1 := strconv.Atoi(v1)
	if err1 != nil {
		panic("Not working i1")
	}

	i2, err2 := strconv.Atoi(v2)
	if err2 != nil {
		panic("Not working i2")
	}

	return i1, i2
}

