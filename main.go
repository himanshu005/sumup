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

## Function #2: Known/supported card schemes

Card Scheme (Visa, MasterCard, JCB, etc) can be detected by the first digits of the card and the length of the card. 

**Example**

| Scheme           | Ranges           | Number of Digits | Example number   |
|---               |---               |---               |---
| American Express | 34,37            | 15               | 378282246310005  |
| JCB              | 3528-3589        | 16-19            | 3530111333300000 |
| Maestro          | 50, 56-58, 6     | 12-19            | 6759649826438453 |
| Visa             | 4                | 13,16,19         | 4012888888881881 |
| MasterCard       | 2221-2720, 51-55 | 16               | 5105105105105100 |


func getSupportedCardSchemes(str string) string {
	str = strings.Join(strings.Fields(str), "")
	num := 0
	for i := 0; i < len(str); i++ {
		s := str[i]
		if s >= 48 && s <= 57 {
			n, _ := strconv.Atoi(string(s))
			l := len(str)
			num = num*10 + n

			if num == 4 && (l == 13 || l == 16 || l == 19) {
				return "Visa"
			} else if (num == 6 || num == 50 || (num >= 56 && num <= 58)) && (l >= 12 && l <= 19) {
				return "Maestro"
			} else if (num == 34 || num == 37) && l == 15 {
				return "American Express"
			} else if ((num >= 51 && num <= 55) || (num >= 2221 && num <= 2720)) && l == 16 {
				return "MasterCard"
			} else if (num >= 3528 && num <= 3589) && (l >= 16 || l <= 19) {
				return "JCB"
			}

		}

	}
	return "none"
}