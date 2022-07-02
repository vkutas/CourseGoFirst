//https://contest.yandex.ru/contest/25606/problems/F/
// F. .сревеР
// Ограничение времени	1 секунда
// Ограничение памяти	64Mb
// Ввод	стандартный ввод или input.txt
// Вывод	стандартный вывод или output.txt
// Напишите программу, которая считывает 4 строки - сообщения, а затем еще одну - имя автора.

// После чего ваша программа выводит 4 введенные строки в обратном порядке вместе с именем автора на каждой строке.

// Формат ввода
// 4 однословные строки - сообщения. И 1 строка - имя автора. Каждая строка вводится с новой строки.
package main

import (
	"fmt"
)

func main() {
	var (
		word1 string
		word2 string
		word3 string
		word4 string
		name  string
	)
	fmt.Scan(&word1)
	fmt.Scan(&word2)
	fmt.Scan(&word3)
	fmt.Scan(&word4)
	fmt.Scan(&name)

	fmt.Println(word4 + " - " + name)
	fmt.Println(word3 + " - " + name)
	fmt.Println(word2 + " - " + name)
	fmt.Println(word1 + " - " + name)

}
