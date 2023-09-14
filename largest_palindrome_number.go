package main

import "fmt"

func isPalindrome(n int) bool {
	reversed := 0
	temp := n
	for temp != 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
	}
	return reversed == n
}
func largestPalindromeProduct(min, max int) (int, int, int) {
	maxproduct := 0
	var m1, m2 int
	for i := max; i >= min; i-- {
		for j := i; j >= min; j-- {
			product := i * j
			if isPalindrome(product) && product > maxproduct {
				maxproduct = product
				m1 = i
				m2 = j
			}
		}
	}
	return maxproduct, m1, m2
}
func main() {
	var min, max int
	fmt.Print("Enter the minimum value: ")
	fmt.Scan(&min)
	fmt.Print("Enter the maximum value: ")
	fmt.Scan(&max)
	result, m1, m2 := largestPalindromeProduct(min, max)
	fmt.Printf("The largest palindrome product between %d and %d is: %d\n", min, max, result)
	fmt.Printf("The multiplicands are: %d and %d\n", m1, m2)
}
