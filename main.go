package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertingResult(n int16) string {
	return convertHelper(n, "")
}

// ^0.6 рекурсивно находим и вычтаем числа получаем римское число и конкатерируем
func convertHelper(n int16, result string) string {
	digits := map[string]int16{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "L": 50, "XC": 90, "C": 100}
	tmpNum := 0
	tmpStr := ""

	if n == 0 {
		return result
	}

	for key, value := range digits {
		if value <= n && value > int16(tmpNum) {
			tmpNum = int(value)
			tmpStr = key
		}
	}

	result += tmpStr
	return convertHelper(n-int16(tmpNum), result)
}

// ^0.5 функция проверки символов и выполнения выражения + - / *
func checkingSign(sing string, x, y float32) float32 {
	switch sing {
	case "+":
		fmt.Println("Результат сложения двух чисел:")
		return x + y
	case "-":
		fmt.Println("Результат разности двух чисел:")
		return x - y
	case "/":
		fmt.Println("Результат деления делимого на делитель:")
		return x / y
	case "*":
		fmt.Println("Результат произведения двух чисел:")
		return x * y
	default:
		fmt.Println(" Калькулятор не умеет выполнять такую операцию")
		return -1
	}
}

// ^ 0.4 Определяет какое число пришло
func convertingRomanNumber(romanNumerals string) float32 {
	switch romanNumerals {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "XI":
		return 9
	case "X":
		return 10
	default:
		panic("Калькулятор должен принимать на вход РИМСКИЕ числа от 1 до 10 включительно, не более.")
	}
}

// ~ 0.3 функция  конвертирует римскими числа в арабские
func checkingRomanNumerals(firstNum, lastNum, sing string) {
	firstConvertNum := convertingRomanNumber(firstNum)
	lastConvertNum := convertingRomanNumber(lastNum)

	if firstConvertNum < lastConvertNum {
		panic("Нельзя вычитать из меньшего числа большее, так как в римской системе нет отрицательных чисел.")
	}
	// получаем результат операции 2-х чисел
	secondNums := checkingSign(sing, firstConvertNum, lastConvertNum)
	//получаем результат преведенного в римскую систему счисления
	result := convertingResult(int16(secondNums))
	fmt.Println(result)

}

// ^ _02 проверка на содержание в строке данных
func checkTypeNumbers(newArr []string) {
	if len(newArr) != 3 {
		panic("Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Повторите попытку снова")
	} else {
		firstValue, lastValue, mathSymbols := newArr[0], newArr[2], newArr[1]
		//Как конвертировать строку в числовой тип ч точкой
		firstNum, _ := strconv.ParseFloat(firstValue, 32)
		lastNum, _ := strconv.ParseFloat(lastValue, 32)

		// приводим данные к булевому типу
		firstTypeNum := firstNum != 0
		lastTypeNum := lastNum != 0
		//Проверяем на то какие два значения приходят
		if firstTypeNum && lastTypeNum {
			// Оба числа
			if firstNum > 10 || lastNum > 10 {
				panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. ")
			} else if math.Mod(float64(firstNum), 1) != 0 || math.Mod(float64(lastNum), 1) != 0 {
				fmt.Println("Калькулятор умеет работать только с целыми числами.")
			} else {
				res := checkingSign(mathSymbols, float32(firstNum), float32(lastNum))
				fmt.Println(res)
			}
		} else if firstTypeNum || lastTypeNum {
			panic("Нельзя использовать одновременно разные системы счисления. И отправлять вместо числа не числа.")
		} else {
			checkingRomanNumerals(firstValue, lastValue, mathSymbols)

		}
	}
}

// ^ _01 Запуск консоли и предача значений
func inputConsole() {
	fmt.Println("Введите арифметическую операцию:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	res := strings.Split(input, " ") //
	checkTypeNumbers(res)
}

func main() {
	inputConsole()
}
