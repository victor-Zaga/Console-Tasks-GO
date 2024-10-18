package main

import (
	"fmt"
)

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func checkNumChetnost() {
	fmt.Printf("Введите число: ")
	var num int
	fmt.Scan(&num)
	var resStr string = evenOrOdd(num)
	if resStr == "even" {
		fmt.Printf("Число четное\n\n")
	} else if resStr == "odd" {
		fmt.Printf("Число нечетное\n\n")
	}
}

func factorialRec(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorialRec(num-1)
}

func factorialCalc() {
	fmt.Print("Введите число: ")
	var number int
	fmt.Scan(&number)
	res := factorialRec(number)
	if res == -1 {
		fmt.Printf("Факториал числа не существует\n\n")
	} else {
		fmt.Printf("Факториал числа %d равен: %d\n\n", number, res)
	}
}

func sumArray(array []int) int {
	var sum int
	for _, elem := range array {
		sum += elem
	}
	return sum
}

func sumArrayElements() {
	fmt.Print("Введите количество элементов массива: ")
	var size int
	fmt.Scan(&size)
	if size <= 0 {
		fmt.Print("Количество должно быть положительным. \n\n")
		return
	}

	var arr []int
	fmt.Println()
	for i := 0; i < size; i++ {
		fmt.Print("Введите ", i+1, " элемент массива: ")
		var elem int
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}

	sum := sumArray(arr)
	fmt.Printf("\nСумма элементов массива: %d\n\n", sum)
}

// меняем перввую и последнюю букву строки, пока индексы не будут равны(четное кол-во букв), либо не зайдут друг за друга(нечетное количество букв)
// завершаем цикл и преобразуем массив runr (int32) в перевернутую строку
func reverseWord(str string) string {
	runes := []rune(str) // []int 32

	for start, end := 0, len(runes)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}
	return string(runes)
}

func reverse() {
	fmt.Print("Введите слово для реверса: ")
	var str string
	fmt.Scan(&str)
	fmt.Printf("Перевернутая строка: %s\n\n", reverseWord(str))
}

func main() {
	// variables

	//var num int // не init
	//var str1 string = "fwef"
	//var str2 = "fwef"
	//str3 := "duasigy"
	//const STR = "HIHI"

	var start bool = true

	for start {
		fmt.Println("\nВыберите задание: ")
		fmt.Println("1. Чётные и нечётные числа.")
		fmt.Println("2. Сумма чисел массива.")
		fmt.Println("3. Факториал числа.")
		fmt.Println("4. Обратная строка.")
		fmt.Println("5. Закончить.")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println()
			checkNumChetnost()
		case 2:
			fmt.Println()
			sumArrayElements()
		case 3:
			fmt.Println()
			factorialCalc()
		case 4:
			fmt.Println()
			reverse()
		case 5:
			fmt.Println()
			fmt.Print("Пока.")
			start = false
		default:
			fmt.Println()
			fmt.Print("Нет такого пункта.\n")
		}
	}
}
