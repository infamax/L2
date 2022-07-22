package main

import "testing"

type test struct {
	name         string
	inputData    string
	expectedRes  string
	expectedFlag bool
}

func TestUnpackString(t *testing.T) {
	tests := []test{
		{
			name:         "incorrect numeric string case",
			inputData:    "45",
			expectedRes:  "",
			expectedFlag: false,
		},
		{
			name:         "empty string case",
			inputData:    "",
			expectedRes:  "",
			expectedFlag: true,
		},
		{
			name:         "string without numbers case",
			inputData:    "abcd",
			expectedRes:  "abcd",
			expectedFlag: true,
		},
		{
			name:         "standard case",
			inputData:    "a4bc2d5e",
			expectedRes:  "aaaabccddddde",
			expectedFlag: true,
		},
		{
			name:         "string start numeric symbol case",
			inputData:    "2a",
			expectedRes:  "",
			expectedFlag: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, flag := UnpackString(tt.inputData)
			if flag != tt.expectedFlag {
				t.Errorf("flag = %v, expectedFlag = %v", flag, tt.expectedFlag)
			}

			if res != tt.expectedRes {
				t.Errorf("res = %v, expectedRes = %s", res, tt.expectedRes)
			}
		})
	}
}
