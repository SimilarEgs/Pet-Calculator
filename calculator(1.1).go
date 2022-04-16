package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var (
		action, function string
		x, y             int
	)

	res := make([]int, 0, 10)

	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) int {
		return x - y
	}

	divideFunc := func(x, y int) int {
		return x / y
	}

	multiplyFunc := func(x, y int) int {
		return x * y
	}

	for {
		fmt.Println("\nChoice action: | + | - | * | / | ? |")
		fmt.Scanln(&action)

		if action == "+" || action == "-" || action == "*" || action == "/" {

			fmt.Print("\nEnter first number: ")

			if _, err := fmt.Scanln(&x); err == nil {

				fmt.Print("\nEnter second number: ")
				fmt.Scanln(&y)

				switch action {
				case "+":
					fmt.Printf("\nAddition result: %d\n", calculate(x, y, sumFunc))
					res = append(res, calculate(x, y, sumFunc))
				case "-":
					fmt.Printf("\nSubtraction result: %d\n", calculate(x, y, subtractFunc))
					res = append(res, calculate(x, y, subtractFunc))
				case "*":
					fmt.Printf("\nMultiplication result: %d\n", calculate(x, y, multiplyFunc))
					res = append(res, calculate(x, y, multiplyFunc))
				case "/":
					fmt.Printf("\"Division result %d\n", calculate(x, y, divideFunc))
					res = append(res, calculate(x, y, divideFunc))
				}
			} else {
				err = errors.New("\nEror: invalid input!\n")
				fmt.Print(err)
			}
		} else if action == "?" {
			fmt.Print("\n1. List of resualts \n2. Back to calculator \n3. Exit\n")
			var exit int
			for exit == 0 {
				fmt.Print("\nNumber of the function: ")
				fmt.Scanln(&function)
				fmt.Println()
				switch function {
				case "1":
					if len(res) >= 1 {
						for i, v := range res {
							fmt.Printf("%d. %d\n", i+1, v)
						}
					} else {
						fmt.Println("Eror: result list is empty")
					}
				case "2":
					exit++
				case "3":
					fmt.Println("Bye!")
					os.Exit(0)
				default:
					fmt.Println("\nEror: invalid function")
				}
			}
		} else {
			fmt.Println("\nEror: not an option")
		}
	}
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}
