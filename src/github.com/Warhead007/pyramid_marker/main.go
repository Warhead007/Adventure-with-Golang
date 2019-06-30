package main

import "fmt"

func main() {
	var asterisk string
	var countLevel int
	var printSpace string
	var input int

	fmt.Print("Enter number of pyramid level : ")
	fmt.Scan(&input)
	fmt.Println("Number of pyramid Level : ", input)

	////loop for create pyramid////
	for i := input; i > 0; i-- {
		////set space to 0////
		printSpace = ""
		////loop for create space////
		for i := input - 1; i > 0; i-- {
			printSpace += " "
		}
		////for correct space number////
		input--

		////first time print only *////
		if countLevel == 0 {
			asterisk = "*"
		} else {
			asterisk += "**"
		}
		////for check if only////
		countLevel++
		////print space and *////
		fmt.Println(printSpace + asterisk)

	}

}
