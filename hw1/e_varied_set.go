// Ограничение времени	1 секунда
// Ограничение памяти	64Mb
// Ввод	стандартный ввод или input.txt
// Вывод	стандартный вывод или output.txt
// Аркадий решил поделиться своим приложением со всем миром! Для этого ему нужно добавить в свое приложение
// интерактивности. Помогите Аркадию.

// Напишите программу, которая принимает на вход 4 строки (гарантируется, что в каждой по одному слову -
// 	нет пробелов): drink, meal, subMeal, time

// И выводит все в формате : "I wanna some '[drink]', my favorite meal - '[meal]'. Give me also '[subMeal]'. I will come soon... '[time]'"

// Важный момент: помните, что строковые переменные внутри print() можно не только перечислять через запятую,
//  но и складывать между собой. Говорят, это называется - конкатенация.

// Формат ввода
// 4 однословные строки. Каждая с новой строчки.

// Формат вывода
// Ответ, требуемый на задачу.
package main

import (
	"fmt"
)

func main() {
	var (
		drink   string
		meal    string
		subMeal string
		time    string
	)
	fmt.Scan(&drink)
	fmt.Scan(&meal)
	fmt.Scan(&subMeal)
	fmt.Scan(&time)
	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'",
		drink, meal, subMeal, time)
}
