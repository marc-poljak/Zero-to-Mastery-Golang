package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {

	quiz1, quiz2, quiz3 := 8, 9, 10

	if quiz1 > quiz2 {
		fmt.Println("quiz1 is greater than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 is less than quiz2")
	} else {
		fmt.Println("quiz1 is equal to quiz2")
	}

	// Calculate the average score
	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("You passed the quiz!")
	} else {
		fmt.Println("You failed the quiz!")
	}

}
