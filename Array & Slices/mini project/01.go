package main

import "fmt"

func findMini(n [5]int) int {
	mini := n[0]
	for i := 1; i < 5; i++ {
		if n[i] < mini {
			mini = n[i]
		}
	}
	return mini
}

func findMax(a [5]int) int {
	max := a[0]
	for i := 1; i < 5; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func findSum(num [5]int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func main() {

	var num [5]int

	fmt.Println("Enter 5 numbers:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&num[i])
	}

	fmt.Println("Your numbers:", num)

	sum := findSum(num)
	fmt.Println("Sum:", sum)

	maxValue := findMax(num)
	fmt.Println("Biggest number:", maxValue)

	minValue := findMini(num)
	fmt.Println("Smallest number:", minValue)
}