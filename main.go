// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "5237 2516 2477 8133"
	fmt.Println(isValidCreditCardNumber(a))

}
/***
	## Function #1: Validity of the card number 

	The following algorithm can be used to check validity of a card number:

	1. Starting from the right, replace each **second** digit of the card number with its doubled value
	2. When doubling a digit produces a 2-digit number (e.g 6 produces 12), then add those 2 digits (1+2 = 3)
	3. Sum up all the digits

	The card number is valid if the sum is divisible by 10

	**Example**: Let's check if `5237 2516 2477 8133` is a valid credit card number.

	1. Double each second digit: **10** 2 **6** 7 **4** 5 **2** 6 **4** 4 **14** 7 **16** 1 **6** 3
	2. Add 2-digit numbers: **1** 2 6 7 4 5 2 6 4 4 **5** 7 **7** 1 6 3
	3. Sum up all the digits: 70

	70 is divisible by 10, so `5237 2516 2477 8133` is a **valid** credit card number

	Please implement a function that given a credit card number returns if it is valid 
*/

func isValidCreditCardNumber(str string) bool {
	count := 0
	sum := 0
	for i := 0; i < len(str); i++ {
		s := str[i]
		if s != 32 {
			n, _ := strconv.Atoi(string(s))
			count++
			if count%2 != 0 && n != 0 {
				n *= 2
				remainder := n % 10
				n = n / 10
				n += remainder
			}
			sum += n
		}
	}
	return sum%10 == 0
}
