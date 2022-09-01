package GoCalculator

import "fmt"

func main() {
	var number1 int
	var number2 int
	var exit bool
	var input int

	fmt.Println("Hello and welcome to the CALCULATOR!!!")
	fmt.Println("Enter two numbers: ")
	fmt.Scan(&number1)
	fmt.Scan(&number2)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Add the numbers")
		fmt.Println("2. Multiply the numbers")
		fmt.Println("3. Divide the numbers")
		fmt.Println("4. Exit the application")

		fmt.Scan(&input)

		if input == 1 {
			fmt.Println("The amount of the two numbers are: %v", helper.addNumbers(number1, number2))
		} else if input == 2 {
			fmt.Println("The multiplication of the numbers are: %v", helper.multiplyNumbers(number1, number2))
		} else if input == 3 {
			fmt.Println("The division of the numbers are: %v", helper.divideNumbers(number1, number2))
		} else if input >= 4 {
			exit = true
			fmt.Println("Bye bye!")
		} else if input < 1 {
			fmt.Println("The value you have entered is too low, smacking your ass for being naughty!")
		}

		if exit == true {
			break
		}
	}
}

func authenticated() bool {
	fmt.Println("Please enter the password!")
	var input string
	fmt.Scan(&input)

	if input == "RobinHood" {
		fmt.Println("Yaay, you have entered the calculator!")
		return true
	} else {
		fmt.Println("Sorry, wrong password!")
		return false
	}
}
