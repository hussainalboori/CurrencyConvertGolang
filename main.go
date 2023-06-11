package main

import "fmt"

func main() {
	// USD
	USDtoEURO := 1.09
	USDtoBHD := 0.38

	// BHD
	BHDtoEURO := 0.41

	var from string
	var to string
	var amount float64

	fmt.Println("Choose main currency ( FROM ): \n")
	fmt.Println("1- USD \n2- BHD \n3- EURO \n")
	fmt.Print("Your input: ")
	fmt.Scanln(&from)
	fmt.Println()

	if from == "4" {
		fmt.Println("###############")
		fmt.Println("Good Bye")
		fmt.Println("###############")
	}

	if from < "1" || from > "4" {
		fmt.Println("Please choose between 1 - 4")
		fmt.Println("___________________________________________________________________________________________________________________________________________________________________________________________________________________")

	}

	for {
		fmt.Println("changing to type ( TO )")
		fmt.Println("1- USD \n2- BHD \n3- EURO \n4- Exit Program \n")
		fmt.Print("Your input: ")
		fmt.Scanln(&to)
		fmt.Println()

		if to < "1" || to > "4" {
			fmt.Println("Please choose between 1 - 3")
			fmt.Println()
			continue
		} else if to == "4" {
			fmt.Println("###############")
			fmt.Println("Good Bye")
			fmt.Println("###############")
			break
		}

		if from == to {
			fmt.Println("Can't be the same currency.")
			fmt.Println()
			continue
		}

		fmt.Print("Amount: ")
		fmt.Scanln(&amount)
		fmt.Println()

		// Check if amount is a valid number greater than 0
		if amount <= 0 {
			fmt.Println("Amount cannot be less than 1.")
			continue
		}

		if from == "1" && to == "2" {
			result := amount * USDtoBHD
			fmt.Printf("%.2f USD is equivalent to %.2f BHD\n", amount, result)

		} else if from == "1" && to == "3" {
			result := amount * USDtoEURO
			fmt.Printf("%.2f USD is equivalent to %.2f EURO\n", amount, result)

		} else if from == "2" && to == "1" {
			result := amount / USDtoBHD
			fmt.Printf("%.2f BHD is equivalent to %.2f USD\n", amount, result)

		} else if from == "2" && to == "3" {
			result := amount / BHDtoEURO
			fmt.Printf("%.2f BHD is equivalent to %.2f EURO\n", amount, result)

		} else if from == "3" && to == "1" {
			result := USDtoEURO * amount
			fmt.Printf("%.2f EURO is equivalent to %.2f USD\n", amount, result)

		} else if from == "3" && to == "2" {
			result := amount * BHDtoEURO
			fmt.Printf("%.2f EURO is equivalent to %.2f BHD\n", amount, result)
		}
		fmt.Println()
		fmt.Println("___________________________________________________________________________________________________________________________________________________________________________________________________________________")
		break
	}
}
