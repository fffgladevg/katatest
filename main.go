package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func romanToArabic(roman string) int {
	roman = strings.ToUpper(roman)
	arabic := 0
	for len(roman) > 0 {
		for key, value := range romanNumerals {
			if strings.HasPrefix(roman, key) {
				arabic += value
				roman = roman[len(key):]
			}
		}
	}
	return arabic
}

func calculate(expression string) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат ввода")
		return
	}

	a, operator, b := parts[0], parts[1], parts[2]
	var isRoman bool

	if _, err := strconv.Atoi(a); err != nil {
		isRoman = true
	}

	var operandA, operandB int

	if isRoman {
		operandA = romanToArabic(a)
		operandB = romanToArabic(b)
	} else {
		var err error
		operandA, err = strconv.Atoi(a)
		if err != nil {
			fmt.Println("Неправильное число:", a)
			return
		}

		operandB, err = strconv.Atoi(b)
		if err != nil {
			fmt.Println("Неправильное число:", b)
			return
		}
	}

	var result int

	switch operator {
	case "+":
		result = operandA + operandB
	case "-":
		result = operandA - operandB
	case "*":
		result = operandA * operandB
	case "/":
		if operandB == 0 {
			fmt.Println("Деление на ноль")
			return
		}
		result = operandA / operandB
	default:
		fmt.Println("Неправильная операция:", operator)
		return
	}

	if isRoman {
		if result <= 0 {
			fmt.Println("Результат меньше 1 для римских чисел")
			return
		}
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 {
		return "Недопустимое значение"
	}

	var result strings.Builder

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (например, 2 + 3): ")
		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSpace(expression)

		if strings.ToLower(expression) == "выход" {
			break
		}

		calculate(expression)
	}
}
