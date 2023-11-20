package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Вывод приветствия
	fmt.Println("~ ~ ~ ~ ~ КАЛЬКУЛЯТОР ~ ~ ~ ~ ~")

	// Правила ввода чисел
	fmt.Println("Вводимые числа должны быть следующего вида:")
	fmt.Println("\n1. Aрабские: - (От 1 до 10 включительно)")
	fmt.Println("Пример ввода: 2 + 2 ")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println("2. Римские: - (От I до X включительно)")
	fmt.Print("Пример ввода: II + II\n")

	// Подсказка о допустимых выражениях
	fmt.Println("____________________________________________________")
	fmt.Println(" F.A.Q Недопустим ввод выражение вида: I + 2, 2 + I;")
	fmt.Println("____________________________________________________")

	// Чтение ввода пользователя
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	// Используйте strings.Fields для разделения ввода на слова (поля)
	fields := strings.Fields(input)
	if len(fields) != 3 {
		fmt.Println("Ошибка: некорректное выражение")
		return
	}

	// Получение чисел и оператора
	a := fields[0]
	operator := fields[1]
	b := fields[2]

	// Проверка на отрицательное число
	isNegative := false
	if strings.HasPrefix(a, "-") {
		isNegative = true
		a = a[1:]
	}

	// Разделение ввода на числа и операцию
	operatorIndex := strings.IndexAny(operator, "+-*/")
	if operatorIndex == -1 {
		fmt.Println("Ошибка: некорректное выражение")
		return
	}

	// Проверка, является ли число арабским, и применение отрицания только для арабских чисел
	if isRomanNumber(a) {
		// Если первое число римское, и оно отрицательное, выводим ошибку
		if isNegative {
			fmt.Println("Ошибка: в римской системе нет отрицательных значений")
			return
		}
	} else {
		var errA error
		a, errA = processArabicNumber(a, isNegative)
		if errA != nil {
			fmt.Println(errA)
			return
		}
	}

	if input[operatorIndex] == '/' && b == "0" {
		fmt.Println("Ошибка: деление на ноль")
		return
	}

	// Проверка диапазона для арабских чисел
	numA, validA := parseNumber(a, isRomanNumber(a), "первое")
	numB, validB := parseNumber(b, isRomanNumber(b), "второе")
	if !validA || !validB || numA < 1 || numA > 10 || numB < 1 || numB > 10 {
		fmt.Println("Ошибка: некорректные числа")
		return
	}

	// Проверка, что оба числа имеют один и тот же тип (или оба римские, или оба арабские)
	if isRomanNumber(a) != isRomanNumber(b) {
		fmt.Println("Ошибка: использование смешанных типов чисел недопустимо")
		return
	}

	// Разрешение операции минус для смешанных типов чисел
	if input[operatorIndex] == '-' && isRomanNumber(a) && isRomanNumber(b) {
		// Ничего не делаем, разрешаем операцию минус для смешанных типов
	} else if isRomanNumber(a) != isRomanNumber(b) {
		fmt.Println("Ошибка: использование смешанных типов чисел недопустимо")
		return
	}

	// Применение отрицания, если число отрицательное
	if isNegative {
		numA *= -1
	}

	// Выполнение операции
	result := 0

	switch operator {
	case "+":
		result = numA + numB

	case "-":
		result = numA - numB

	case "*":
		result = numA * numB

	case "/":
		result = numA / numB

	default:
		fmt.Println("Ошибка: некорректная операция")
		return
	}

	// Вывод результата
	isRomanResult := isRomanNumber(a) && isRomanNumber(b)
	if isRomanResult {
		romanResult := arabicToRoman(result)
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}

// Функция для проверки, являются ли числа римскими
func isRomanNumber(str string) bool {
	romanDigits := "IVX"
	for _, char := range str {
		if !strings.ContainsAny(string(char), romanDigits) {
			return false
		}
	}
	return true
}

// Функция для обработки отрицательных чисел и преобразования арабских чисел
func processArabicNumber(input string, isNegative bool) (string, error) {
	if isNegative {
		numA, err := strconv.Atoi(input)
		if err != nil || numA < 1 || numA > 10 {
			return input, fmt.Errorf("Ошибка: введите отрицательное арабское число от 1 до 10 включительно")
		}
		// Вернуть отрицательное число в виде строки
		return strconv.Itoa(numA), nil
	}
	// Вернуть положительное число без изменений
	return input, nil
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

func parseNumber(str string, isRoman bool, variableName string) (int, bool) {
	if isRoman {
		romanToArabic := map[string]int{
			"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
			"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		}

		// Проверка, что введенное римское число существует в мапе
		num, ok := romanToArabic[str]
		if !ok {
			fmt.Printf("Ошибка: введите %s римское число от I до X включительно\n", variableName)
			return 0, false
		}
		return num, true
	}

	num, err := strconv.Atoi(str)
	if err != nil || num < 1 || num > 10 {
		fmt.Printf("Ошибка: введите %s арабское число от 1 до 10 включительно\n", variableName)
		return 0, false
	}

	return num, true
}
