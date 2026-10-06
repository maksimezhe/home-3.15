package main

import (
	"errors"
	"fmt"
	"math"
)

// 2
func divide(a int, b int) (int, interface{}) {
	if b == 0 {
		return 0, "Ошибка: деление на ноль невозможно"
	}
	return a / b, nil
}

// 3
func checkAge(age int) error {
	if age < 0 || age > 120 {
		return errors.New("указан некорректный возраст")
	}
	return nil
}

// 4
func validatePassword(password string) (bool, error) {
	if len(password) < 6 {
		return false, errors.New("пароль слишком короткий")
	}
	return true, nil
}

// 5
type invalidRadiusError struct{}

func (e invalidRadiusError) Error() string {
	return "радиус не может быть отрицательным"
}

// 6
func calculateCircleArea(radius int) (float64, error) {
	if radius < 0 {
		return 0, invalidRadiusError{}
	}
	area := math.Pi * float64(radius) * float64(radius)
	return area, nil
}

// 7
func findUser(user []string, name string) (int, error) {
	for index, user := range user {
		if user == name {
			return index, nil
		}
	}
	return -1, errors.New("пользователь не найден")
}

func main() {
	//Task 1
	prices := map[string]int{"Хлеб": 190, "Молоко": 120}
	productName := "Хлеб"

	price, ok := prices[productName]

	if ok {
		fmt.Println("Цена:", price)
	} else {
		fmt.Println("Товар отсутствует")
	}

	//Task 2
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", result)
	}
	result2, err2 := divide(10, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Результат:", result2)
	}

	// Task 3
	err = checkAge(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Возраст корректен")
	}

	// Task 4
	valid, err := validatePassword("123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Пароль корректен:", valid)
	}

	valid2, err2 := validatePassword("1234567")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Пароль корректен:", valid2)
	}

	// Task 5
	radiusErr := invalidRadiusError{}
	fmt.Println(radiusErr.Error())

	// Task 6
	area, err := calculateCircleArea(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Площадь круга:", area)
	}

	// Task 7
	users := []string{"Максим", "Бекболат", "Петя"}
	index, err := findUser(users, "Бекболат")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Индекс пользователя:", index)
	}
}
