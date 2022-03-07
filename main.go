package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "5237 2516 2477 8133"
	fmt.Println(isValidCreditCardNumber(a))

}
/***
	## Function #1: Validity of the card number 
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
/***
## Function #2: Known/supported card schemes
*/
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