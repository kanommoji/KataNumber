package kata

import "testing"

func Test_Kata_Input_1_Should_be_kata_1(t *testing.T) {
	expectedResult := "|\n|"

	actualResult := Kata(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", actualResult, expectedResult)
	}
}
