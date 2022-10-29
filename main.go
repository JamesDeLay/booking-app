package main

import (
	"fmt"
)

// Recursive fib
// func fib(num int) int {
// 	if num <= 2 {
// 		return 1
// 	}
// 	return num + fib(num-1)
// }

// Recursive Palindrome
// func isPalindrome(str string) bool {
// 	if str[0] != str[len(str)-1] {
// 		return false
// 	}
// 	if len(str) == 1 {
// 		return true
// 	} else {
// 		return isPalindrome(str[1 : len(str)-1])
// 	}
// }

func main() {
	const conferenceName = "Jimbo's Go Conf"
	const availableTickets = 50
	var remainingTickets = 37

	fmt.Println("Hello world, welcome to", conferenceName, "booking application!!!")
	fmt.Printf("Get your tickets here, %v available, %v left!\n", remainingTickets, availableTickets)

	var (
		userName       string
		email          string
		desiredTickets string
	)

	fmt.Println("Enter username...")
	fmt.Scan(&userName)

	fmt.Println("Enter email...")
	fmt.Scan(&email)

	fmt.Println("Enter number of desired tickets...")
	fmt.Scan(&desiredTickets)

	fmt.Printf("Thank you %v, you've successfully booked %v tickets. A confirmation email will be sent to %v\n", userName, desiredTickets, email)
	remainingTickets -= 1
	fmt.Println(remainingTickets)
}
