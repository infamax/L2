package main

import "testing"

type test struct {
	name        string
	inputData   string
	escapeFlag  bool
	expectedRes string
	expectedErr error
}

func TestUnpackStringAdvanced(t *testing.T) {
	tests := []test{
		{
			name:        "incorrect numeric string case",
			inputData:   "45",
			escapeFlag:  noEscapeSeq,
			expectedRes: "",
			expectedErr: numericFirstSymbolErr,
		},
		{
			name:        "empty string case",
			inputData:   "",
			escapeFlag:  noEscapeSeq,
			expectedRes: "",
			expectedErr: nil,
		},
		{
			name:        "string without numbers case",
			inputData:   "abcd",
			escapeFlag:  noEscapeSeq,
			expectedRes: "abcd",
			expectedErr: nil,
		},
		{
			name:        "standard case",
			inputData:   "a4bc2d5e",
			escapeFlag:  noEscapeSeq,
			expectedRes: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			name:        "string start numeric symbol case",
			inputData:   "2a",
			escapeFlag:  noEscapeSeq,
			expectedRes: "",
			expectedErr: numericFirstSymbolErr,
		},
		{
			name:        "string without pack",
			inputData:   "qwe\\4\\5",
			escapeFlag:  escapeSeq,
			expectedRes: "qwe45",
			expectedErr: nil,
		},
		{
			name:        "pack string escape seq case",
			inputData:   "qwe\\45",
			escapeFlag:  escapeSeq,
			expectedRes: "qwe44444",
			expectedErr: nil,
		},
		{
			name:        "string with escape pack case",
			inputData:   "qwe\\\\5",
			escapeFlag:  escapeSeq,
			expectedRes: "qwe\\\\\\\\\\",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := UnpackStringAdvanced(tt.inputData, tt.escapeFlag)
			if err != tt.expectedErr {
				t.Errorf("err = %v, expectedErr = %v", err, tt.expectedErr)
			}

			if res != tt.expectedRes {
				t.Errorf("res = %v, expectedRes = %s", res, tt.expectedRes)
			}
		})
	}
}
