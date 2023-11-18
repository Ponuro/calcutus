package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Вывод приветствия
	fmt.Println("~ ~ ~ ~ ~ КАЛЬКУЛЯТОР ~ ~ ~ ~ ~")

	// Правила ввода чисел
	fmt.Println("Вводимые числа должны быть следующего вида:")
	fmt.Println("\n1. Aрабские: - (От 1 до 10 включительно)")
	fmt.Println("Пример ввода: 2+2")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println("2. Римские: - (От I до X включительно)")
	fmt.Println("Пример ввода: II+II\n")

	// Подсказка о допустимых выражениях
	fmt.Println("________________________________________________")
	fmt.Println(" F.A.Q Недопустим ввод выражение вида: I+2, 2+I; ")
	fmt.Println("________________________________________________")

	// Чтение ввода пользователя
	var input string
	fmt.Print("\nВведите выражение: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	// Проверка на отрицательное число
	isNegative := false
	if strings.HasPrefix(input, "-") {
		isNegative = true
		input = input[1:]
	}

	// Разделение ввода на числа и операцию
	operatorIndex := strings.IndexAny(input, "+-*/")
	if operatorIndex == -1 {
		fmt.Println("Ошибка: некорректное выражение")
		return
	}

	// Получение чисел и оператора
	a := input[:operatorIndex]
	b := input[operatorIndex+1:]

	// Проверка, являются ли числа римскими
	isRoman := isRomanNumber(a) && isRomanNumber(b)

	// Парсинг чисел
	numA, validA := parseNumber(a, isRoman, "первое")
	numB, validB := parseNumber(b, isRoman, "второе")

	// Проверка деления на ноль
	if operatorIndex < len(input)-1 && input[operatorIndex+1] == '0' && (input[operatorIndex] == '/' || input[operatorIndex] == '%') {
		fmt.Println("Ошибка: деление на ноль")
		return
	}

	// Проверка диапазона для арабских чисел
	if !validA || !validB || numA < 1 || numA > 10 || numB < 1 || numB > 10 {
		fmt.Println("Ошибка: некорректные числа")
		return
	}

	// Применение отрицания, если число отрицательное
	if isNegative {
		numA *= -1
	}

	// Определение операции
	operator := string(input[operatorIndex])

	// Выполнение операции
	result := 0
	switch operator {
	case "+":
		result = numA + numB

	case "-":
		if isRoman && numA < numB {
			fmt.Println("Ошибка: некорректное выражение")
			return
		}
		result = numA - numB

	case "*":
		if isRoman && numA <= numB {
			fmt.Println("Ошибка: некорректное выражение")
			return
		}
		result = numA * numB

	case "/":
		if numB == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = numA / numB

	default:
		fmt.Println("Ошибка: некорректная операция")
		return
	}

	// Вывод результата
	if isRoman {
		romanResult := arabicToRoman(result)
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}

// Функция для проверки, является ли число римским
func isRomanNumber(str string) bool {
	romanDigits := "IVX"
	for _, char := range str {
		if !strings.ContainsAny(string(char), romanDigits) {
			return false
		}
	}
	return true
}

// Функция для преобразования арабского числа в римское
func arabicToRoman(num int) string {
	romanDigits := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	romanValues := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}

	result := ""
	for i := len(romanDigits) - 1; i >= 0; i-- {
		for num >= romanValues[i] {
			result += romanDigits[i]
			num -= romanValues[i]
		}
	}
	return result
}

// Функция для парсинга чисел
func parseNumber(str string, isRoman bool, variableName string) (int, bool) {
	var num int

	if isRoman {
		romanToArabic := map[string]int{
			"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
			"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		}

		// Проверка, что введенное римское число существует в мапе
		if arabicNum, exists := romanToArabic[str]; exists {
			num = arabicNum
		} else {
			if variableName == "первое" {
				fmt.Printf("Ошибка: введите %s римское число от I до X включительно\n", variableName)
				return 0, false
			} else if variableName == "второе" {
				fmt.Printf("Ошибка: введите %s римское число от I до X включительно\n", variableName)
				return 0, false
			}
		}

	} else {
		// Проверка диапазона для арабских чисел
		arabic, err := strconv.Atoi(str)
		if err != nil || arabic < 1 || arabic > 10 {
			fmt.Printf("Ошибка: введите %s арабское число от 1 до 10 включительно\n", variableName)
			return 0, false
		}
		num = arabic
	}

	return num, true
}
