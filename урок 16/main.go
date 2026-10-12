
package main

import "fmt"

func main() {
	// Задание 1
	toolUsage := map[string]int{
		"Go":     3,
		"VSCode": 5,
		"Git":    2,
	}

	for tool, count := range toolUsage {
		fmt.Println(tool, count)
	}

	// Задание 2
	buildStatus := map[string]bool{
		"build": true,
		"run":   false,
	}

	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}

	// Задание 3
	var name string

	fmt.Println("Введите имя:")
	fmt.Scan(&name)

	userInfo := map[string]string{
		"name":       name,
		"isLoggedIn": "true",
	}

	fmt.Println("Пользователь", userInfo["name"], "вошёл в систему")

	// Задание 4
	cpuLoad := map[int]int{
		1: 40,
		2: 65,
		3: 30,
	}

	maxLoad := -1
	core := 0

	for number, load := range cpuLoad {
		if load > maxLoad {
			maxLoad = load
			core = number
		}
	}

	fmt.Println("Ядро с максимальной загрузкой:", core)

	// Задание 5
	examResults := map[string]int{
		"Aruzhan": 85,
		"Dias":    92,
		"Alina":   78,
	}

	for student, grade := range examResults {
		if grade >= 80 {
			fmt.Println(student, "сдал экзамен")
		} else {
			fmt.Println(student, "не сдал экзамен")
		}
	}

	// Задание 6
	words := []string{"go", "is", "fast"}
	wordLength := make(map[string]int)

	for _, word := range words {
		wordLength[word] = len(word)
	}

	fmt.Println(wordLength)

	// Задание 7
	menu := map[string]int{
		"Burger": 2500,
		"Pizza":  3200,
		"Tea":    500,
	}

	var dish string

	fmt.Println("Введите название блюда:")
	fmt.Scan(&dish)

	if price, exists := menu[dish]; exists {
		fmt.Println("Цена:", price)
	} else {
		fmt.Println("Блюдо не найдено")
	}

	// Задание 8
	loginAttempts := map[string]int{
		"admin": 2,
		"guest": 0,
	}

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}

	// Задание 9
	sales := [2][3]int{
		{10, 15, 20},
		{12, 18, 25},
	}

	total := make(map[int]int)

	for shop := 0; shop < 2; shop++ {
		total[shop+1] = 0

		for day := 0; day < 3; day++ {
			total[shop+1] += sales[shop][day]
		}
	}

	fmt.Println(total)

	// Задание 10
	numbers := []int{4, 7, 2, 9, 5}
	numberStatus := make(map[int]string)

	for _, number := range numbers {
		if number%2 == 0 {
			numberStatus[number] = "even"
		} else {
			numberStatus[number] = "odd"
		}
	}

	fmt.Println(numberStatus)

	// Задание 11
	defaultConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	currentConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	// Первое сравнение
	equal := true

	for key, value := range defaultConfig {
		if currentConfig[key] != value {
			equal = false
		}
	}

	if len(defaultConfig) != len(currentConfig) {
		equal = false
	}

	if equal {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	// Изменение конфигурации
	currentConfig["mode"] = "debug"

	equal = true

	for key, value := range defaultConfig {
		if currentConfig[key] != value {
			equal = false
		}
	}

	if len(defaultConfig) != len(currentConfig) {
		equal = false
	}

	if equal {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}