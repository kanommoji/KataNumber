package kata

import "fmt"

func ToDigitalNumbers(number int) string {
	var lcdNumber [8]int
	var kataNumber string
	if number == 1 {
		lcdNumber[4] = 1
		lcdNumber[7] = 1
	}
	if number == 2 {
		lcdNumber[1] = 1
		lcdNumber[3] = 1
		lcdNumber[4] = 1
		lcdNumber[5] = 1
		lcdNumber[6] = 1

	}
	if number == 3 {
		lcdNumber[1] = 1
		lcdNumber[3] = 1
		lcdNumber[4] = 1
		lcdNumber[6] = 1
		lcdNumber[7] = 1

	}

	for i := 0; i < 8; i++ {
		if lcdNumber[i] == 1 {
			if i == 1 || i == 3 || i == 6 {
				kataNumber = fmt.Sprint(kataNumber, "_")
			} else {
				kataNumber = fmt.Sprint(kataNumber, "|")
			}
			if i == 3 && lcdNumber[4] == 0 {
				kataNumber = fmt.Sprint(kataNumber, "\n")
			}
			if i == 1 || i == 4 {
				kataNumber = fmt.Sprint(kataNumber, "\n")
			}
		}
	}

	return kataNumber
}
