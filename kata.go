package kata

import "fmt"

func ToDigitalNumbers(number int) string {
	Hight := []string{" _", " ", " _"}
	Meium := []string{"|_|", " |", " _|"}
	Low := []string{"|_|", " |", "|_"}
	var digitalNumber string

	digitalNumber = fmt.Sprint(digitalNumber, Hight[number], "\n")
	digitalNumber = fmt.Sprint(digitalNumber, Meium[number], "\n")
	digitalNumber = fmt.Sprint(digitalNumber, Low[number])
	return digitalNumber
}
