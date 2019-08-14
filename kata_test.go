package kata

import "testing"

func Test_Kata_Input_1_Should_be_Digital_Number_0(t *testing.T) {
	expectedResult := ` _
|_|
|_|`

	actualResult := ToDigitalNumbers(0)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
func Test_Kata_Input_1_Should_be_Digital_Number_1(t *testing.T) {
	expectedResult := ` 
 |
 |`

	actualResult := ToDigitalNumbers(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Kata_Input_2_Should_be_Digital_Number_2(t *testing.T) {
	expectedResult := ` _
 _|
|_`

	actualResult := ToDigitalNumbers(2)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
