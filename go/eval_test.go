package main

import (
	"testing"
)

type EvalTest struct {
	ExpectedResult int
	Input          string
}

func TestEvaluation(t *testing.T) {
	tests := []EvalTest{
		{
			ExpectedResult: 39,
			Input:          "3 5 8 + *",
		},
		{
			ExpectedResult: 5,
			Input:          "5 3 + 4 / 4 1 / 12 + * 14 +",
		},
	}

	for i, v := range tests {
		result := evaluatePostfix(v.Input)

		if v.ExpectedResult != result {
			t.Fatalf("Case [%d]: result %d != expected %d", i, result, v.ExpectedResult)
		}
	}
}
