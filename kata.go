package kata

import "fmt"

func ToDigitalNumbers(number int) string {
	Hight := []string{` `}
	Meium := []string{` |`}
	Low := []string{` |`}
	var digitalNumber string
	digitalNumber = fmt.Sprint(digitalNumber, Hight[0], "\n")
	digitalNumber = fmt.Sprint(digitalNumber, Meium[0], "\n")
	digitalNumber = fmt.Sprint(digitalNumber, Low[0])
	return digitalNumber
}
