package main

import "fmt"

func main() {

	var price float64
	var discount float64
	var finalPrice float64
	var productName string
	var contProcess string

	for {

		fmt.Print("Enter the product name: ")
		fmt.Scan(&productName)
		price, discount = get_value()
		finalPrice = calculate(price, discount)
		print_text(finalPrice, productName)
		fmt.Print("Do you want to continue (y/n): ")
		fmt.Scan(&contProcess)
		if contProcess != "y" {
			fmt.Println("! See You !")
			break
		}
	}
}

func get_value() (float64, float64) {

	var price float64
	var discount float64
	for {
		fmt.Print("Enter the price: ")
		fmt.Scan(&price)
		if price > 0 {
			break
		}
		fmt.Println("Invalid price")
	}
	for {
		fmt.Print("Enter the discount: ")
		fmt.Scan(&discount)
		if discount >= 0 {
			break
		}
		fmt.Println("Invalid discount")
	}

	return price, discount
}

func calculate(price float64, discount float64) float64 {
	return price - ((discount * price) / 100)
}

func print_text(result float64, name string) {
	printText := fmt.Sprintf("Final price of %s is : %.1f", name, result)
	fmt.Println(printText)
}
