package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxNumber = 10
	minNumber = 1
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите выражение (например, 5 + 3 или V - II): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Разделение входной строки на числа и операцию
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный формат ввода!")
			continue
		}

		num1 := parts[0]
		operator := parts[1]
		num2 := parts[2]

		// Определение, используются ли римские или арабские цифры
		isRoman1 := isRomanNumeral(num1)
		isRoman2 := isRomanNumeral(num2)
		if isRoman1 != isRoman2 {
			fmt.Println("Введены числа разных типов (римские и арабские)")
			continue
		}

		var result interface{}
		if isRoman1 {
			a := romanToInt(num1)
			b := romanToInt(num2)
			if a < minNumber || a > maxNumber || b < minNumber || b > maxNumber {
				fmt.Println("Число вне диапазона от 1 до 10")
				continue
			}
			result = calculateRoman(a, b, operator)
		} else {
			a, _ := strconv.Atoi(num1)
			b, _ := strconv.Atoi(num2)
			if a < minNumber || a > maxNumber || b < minNumber || b > maxNumber {
				fmt.Println("Число вне диапазона от 1 до 10")
				continue
			}
			result = calculateArabic(a, b, operator)
		}

		fmt.Println("Результат:", result)
	}
}

func calculateArabic(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("Неизвестная операция")
		os.Exit(1)
		return 0 // Added return statement to satisfy the compiler
	}
}

func calculateRoman(a, b int, op string) string {
	switch op {
	case "+":
		return intToRoman(a + b)
	case "-":
		if a-b < 1 {
			fmt.Println("Результат меньше 1, что недопустимо для римских чисел")
			os.Exit(1)
		}
		return intToRoman(a - b)
	case "*":
		return intToRoman(a * b)
	case "/":
		return intToRoman(a / b)
	default:
		fmt.Println("Неизвестная операция")
		os.Exit(1)
		return "" // Added return statement to satisfy the compiler
	}
}

func isRomanNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && romanValues[s[i]] > romanValues[s[i-1]] {
			result += romanValues[s[i]] - 2*romanValues[s[i-1]]
		} else {
			result += romanValues[s[i]]
		}
	}
	return result
}

func intToRoman(num int) string {
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i, value := range romanValues {
		count := num / value
		result.WriteString(strings.Repeat(romanSymbols[i], count))
		num -= value * count
	}
	return result.String()
}
