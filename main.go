package main

import "fmt"

func main() {
	name := "waiwai"

	for i := 0; i < 100000; i++ {
		fmt.Println(name, "=>", i)
	}

	zxx := onCal(3, 10)
	fmt.Println(zxx)

	fmt.Println("-->>  fun finish")

	fmt.Println("-->>  feature sum")

	fmt.Println("-->>  feature develop")
}

func onCal(number1, number2 int) int {
	return number1 + number2
}
