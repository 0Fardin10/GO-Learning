package main

import "fmt"

func findSum(num [5]int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
func 
func main() {
	var num [5]int
	fmt.Println("Enter 5 number :")
	for i := 0; i < 5; i++ {
		fmt.Scan(&num[i])
	}
	fmt.Println("Your number is", num)
	result := findSum(num)
	fmt.Println("Here the sum is :", result)
}
