package main

import "testing"

type test struct {
	name        string
	inputData   string
	expectedRes string
	expectedErr error
}

func TestUnpackString(t *testing.T) {
	tests := []test{
		{
			name:        "incorrect numeric string case",
			inputData:   "45",
			expectedRes: "",
			expectedErr: numericFirstSymbolErr,
		},
		{
			name:        "empty string case",
			inputData:   "",
			expectedRes: "",
			expectedErr: nil,
		},
		{
			name:        "string without numbers case",
			inputData:   "abcd",
			expectedRes: "abcd",
			expectedErr: nil,
		},
		{
			name:        "standard case",
			inputData:   "a4bc2d5e",
			expectedRes: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			name:        "string start numeric symbol case",
			inputData:   "2a",
			expectedRes: "",
			expectedErr: numericFirstSymbolErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UnpackString(tt.inputData)
			if err != tt.expectedErr {
				t.Errorf("err = %v, expectedErr = %v", err, tt.expectedErr)
			}

			if res != tt.expectedRes {
				t.Errorf("res = %v, expectedRes = %s", res, tt.expectedRes)
			}
		})
	}
}
