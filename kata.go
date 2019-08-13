package kata

import "fmt"

func Kata(number int) string {
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

	if lcdNumber[1] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "_\n")
		lcdNumber[1] = 0
	}
	if lcdNumber[3] == 1 && lcdNumber[4] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "_|\n")
		lcdNumber[3] = 0
		lcdNumber[4] = 0
	}
	if lcdNumber[5] == 1 && lcdNumber[6] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "|_\n")
		lcdNumber[5] = 0
		lcdNumber[6] = 0
	}
	if lcdNumber[6] == 1 && lcdNumber[7] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "_|\n")
		lcdNumber[6] = 0
		lcdNumber[7] = 0
	}

	if lcdNumber[4] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "|\n")
		lcdNumber[4] = 0
	}
	if lcdNumber[7] == 1 {
		kataNumber = fmt.Sprint(kataNumber, "|\n")
		lcdNumber[7] = 0
	}

	return kataNumber
}
