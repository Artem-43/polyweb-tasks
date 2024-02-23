package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (число операция число): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			return
		}

		// Удаление символа новой строки из ввода
		input = strings.TrimSpace(input)

		// Проверка на выход из программы
		if input == "exit" {
			fmt.Println("Программа завершена.")
			return
		}

		// Разделение ввода на число1, операцию и число2
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Некорректный ввод. Ожидается: Число Операция Число")
			continue
		}

		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования первого числа:", err)
			continue
		}

		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования второго числа:", err)
			continue
		}

		var result float64
		switch parts[1] {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Деление на ноль невозможно.")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("Некорректная операция. Допустимые операции: +, -, *, /")
			continue
		}

		fmt.Printf("Результат: %.2f\n", result)
	}
}
