package main

import (
	"errors"
	"fmt"
	"strconv"
)

func p(a, b int, s string) (int, error) {
	switch {
	case s == "+":
		return a + b, nil
	case s == "-":
		return a - b, nil
	case s == "*":
		return a * b, nil
	case s == "/":
		return a / b, nil
	default:
		return 0, errors.New("invalid sign")
	}
}

func calculate(g, h, sign string) (string, error) {
	k := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	v := k[g]
	m := k[h]
	if v != 0 && m != 0 {
		result, err := p(v, m, sign)
		if err != nil {
			return "", err
		}

		if result < 1 {
			return "", errors.New("roman number is less than zero")
		}

		return intToRoman(result), nil
	}

	if v != 0 && m == 0 {
		return "", errors.New("invalid value")
	}

	if v == 0 && m != 0 {
		return "", errors.New("invalid value")
	}

	e1, err := strconv.Atoi(g)
	if err != nil {
		return "", err
	}

	e2, err := strconv.Atoi(h)
	if err != nil {
		return "", err
	}

	err = validate(e1, e2)
	if err != nil {
		return "", err
	}

	result, err := p(e1, e2, sign)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", result), nil
}

func validate(a, b int) error {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		return nil
	}

	return errors.New("invalid values")
}

func intToRoman(a int) string {
	ints := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i := 0; i < len(ints); i++ {
		for a >= ints[i] {
			a -= ints[i]
			result += romans[i]
		}
	}

	return result
}

func main() {
	var a1, s, b1 string
	_, err := fmt.Scanf("%s %s %s", &a1, &s, &b1)
	if err != nil {
		panic(err)
	}

	result, err := calculate(a1, b1, s)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
