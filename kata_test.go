package kata

import "testing"

func Test_Kata_Input_1_Should_be_kata_1(t *testing.T) {
	expectedResult := "|\n|\n"

	actualResult := Kata(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Kata_Input_2_Should_be_kata_2(t *testing.T) {
	expectedResult := "_\n_|\n|_\n"

	actualResult := Kata(2)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

func Test_Kata_Input_3_Should_be_kata_3(t *testing.T) {
	expectedResult := "_\n_|\n_|\n"

	actualResult := Kata(3)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}
