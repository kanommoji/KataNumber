package kata

import "testing"

func Test_Kata_Input_1_Should_be_kata_1(t *testing.T) {
	expectedResult := "|\n|"

	actualResult := ToDigitalNumbers(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Kata_Input_2_Should_be_kata_2(t *testing.T) {
	expectedResult := "_\n_|\n|_"

	actualResult := ToDigitalNumbers(2)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Kata_Input_3_Should_be_kata_3(t *testing.T) {
	expectedResult := "_\n_|\n_|"

	actualResult := ToDigitalNumbers(3)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
