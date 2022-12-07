package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "*", "/"}

var arabic = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

var roman = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var txt string
	if sc.Scan() {
		txt = sc.Text()
	}
	arr := strings.Split(strings.TrimSpace(txt), " ")
	res := getSolution(arr)
	fmt.Println(res)
}

// Подготовка данных
func prepareData(temp []string) (string, any, any) {
	correctOperator := false
	if len(temp) != 3 {
		errorObserver("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	for _, v := range operators {
		if v == temp[1] {
			correctOperator = true
			break
		}
	}
	if correctOperator == false {
		errorObserver("недоступная операция [ " + temp[1] + " ]")
	}
	dat1, err1 := strconv.Atoi(temp[0])
	dat2, err2 := strconv.Atoi(temp[2])

	if reflect.TypeOf(err1) != reflect.TypeOf(err2) {
		errorObserver("используются одновременно разные системы счисления")
	}
	if err1 == nil {

		return "int", dat1, dat2
	}
	return "roman", temp[0], temp[2]

}

func checkData(data, data2 any, dataType string) (int, int) {
	switch dataType {
	case "roman":
		d, d2 := romanToInt(data.(string)), romanToInt(data2.(string))
		if d > 10 || d < 1 || d2 > 10 || d2 < 1 {
			errorObserver("числа вне допустимого диапазона ( 1 - 10)")
		}
		return d, d2

	case "int":
		if data.(int) > 10 || data.(int) < 1 || data2.(int) > 10 || data2.(int) < 1 {
			errorObserver("числа вне допустимого диапазона ( 1 - 10)")
		}
		return data.(int), data2.(int)

	}
	return 0, 0
}

// Вычисляет итоговый результат
func getSolution(temp []string) (result any) {
	dateType, date, date2 := prepareData(temp)

	d, d2 := checkData(date, date2, dateType)

	result = calculate(d, d2, temp[1])
	if dateType == "roman" {
		if result.(int) < 1 {
			errorObserver("минимальное значение I")
		}
		result = intToRoman(result.(int))
	}
	return result
}

// Перевод римских чисел в арабские ( V -> 5 )
func romanToInt(romanString string) (result int) {

	text := strings.ToUpper(romanString)
	result = 0

	for i := 0; i < len(arabic); i++ {

		posit := 0
		n := len(arabic) - 1

		for n >= 0 && posit < len(text) {

			if len(roman[n]) <= len(text) && text[posit:len(roman[n])] == roman[n] {
				result += arabic[n]
				text = text[len(roman[n]):]
				if len(text) < 1 || text == "" {
					break
				}
				posit++
			} else {
				n--
			}
		}
	}
	return result
}

// Перевод из абарбских в римские ( 5 -> V )
func intToRoman(number int) (result string) {

	if number >= 4000 || number <= 0 {
		return result
	}
	result = ""
	for i := len(arabic) - 1; i >= 0; i-- {
		for number >= arabic[i] {
			number -= arabic[i]
			result += roman[i]
		}
	}
	return result
}

// Вычисление выражения ( значение1 - 3, значение2 - 3, операция - `+`, результат = 6 и тд ...)
func calculate(firstNumber, secondNumber int, operator string) int {

	defer func() {
		if err := recover(); err != nil {
			errorObserver("divide by zero")
		}
	}()

	switch operator {
	case "+":
		return firstNumber + secondNumber
	case "-":
		return firstNumber - secondNumber
	case "*":
		return firstNumber * secondNumber
	case "/":
		return firstNumber / secondNumber
	default:
		return 0
	}
}

// Обработчик ошибок
func errorObserver(message string) {

	fmt.Println("Error :", message)
	os.Exit(0)
}
