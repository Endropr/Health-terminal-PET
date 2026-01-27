package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\033[34m=== ДОБРО ПОЖАЛОВАТЬ В HEALTH TERMINAL ❘❘ WELCOME TO THE HEALTH TERMINAL ===\033[0m")

	for {
		var weight, height float64

		// Вес
		fmt.Print("\033[33mВведите ваш вес (кг):\033[0m ❘❘ \033[33mEnter your weight (kg):\033[0m ")
		_, err := fmt.Scan(&weight)

		// Косяк юзера
		if err != nil {
			fmt.Println("\033[31mОшибка: введите ЧИСЛО!\033[0m || \033[31mError: enter a NUMBER!\033[0m")
			var dump string
			fmt.Scan(&dump)
			continue
		}

		// Рост
		fmt.Print("\033[33mВведите ваш рост:\033[0m ❘❘ \033[33mEnter your height:\033[0m")
		fmt.Scan(&height)
		if height > 3 {
			height = height / 100
		}

		// Индекс
		fmt.Println("-------------------------------")
		bmi := weight / (height * height)
		fmt.Printf("Ваш индекс массы тела: ❘❘ Your body mass index: \033[44m%.2f\033[0m\n", bmi)

		// ИМТ
		const (
			MinNormalBMI = 18.5
			MaxNormalBMI = 24.9
		)

		// Условия
		if bmi >= MinNormalBMI && bmi <= MaxNormalBMI {
			fmt.Println("\033[32mТвоё тело в полном порядке.\033[0m ❘❘ \033[32mYour body is in perfect order.\033[0m")
		} else if bmi >= MaxNormalBMI {
			fmt.Println("\033[31mКажется, у тебя проблемы с лишним весом, срочно иди в спортзал!\033[0m ❘❘ \033[31mIt seems like you have problems with excess weight, go to the gym immediately!\033[0m")
		} else if bmi < MinNormalBMI {
			fmt.Println("\033[31mВаш ИМТ очень низкий! Вам нужно больше питаться!\033[0m ❘❘ \033[31mYour BMI is really low! You need to eat more!\033[0m")
		}
		// Выход
		fmt.Println("-------------------------------")
		fmt.Print("Нажмите '\033[34mQ\033[0m' для выхода или любую клавишу для повтора: ❘❘ Press '\033[34mQ\033[0m' to exit or any key to repeat: ")
		var userChoice string
		fmt.Scan(&userChoice)

		if strings.ToLower(userChoice) == "q" {
			fmt.Println("\033[33mПока! Здоровье не купишь - его разум дарит!\033[0m ❘❘ \033[33mBye! You can't buy health - your mind gives it!\033[0m")
			break
		}
	}
}
