package main

import (
	"reflect"
	"testing"
)

type test struct {
	name        string
	inputData   []string
	expectedRes map[string][]string
}

func TestGroupAnagrams(t *testing.T) {
	tests := []test{
		{
			name:        "empty slice case",
			inputData:   []string{},
			expectedRes: nil,
		},
		{
			name: "standard case",
			inputData: []string{"пятак", "пятка", "тяпка",
				"листок", "слиток", "столик"},
			expectedRes: map[string][]string{
				"листок": {"слиток", "столик"},
				"пятак":  {"пятка", "тяпка"},
			},
		},
		{
			name:        "one elem slice case",
			inputData:   []string{"пятак"},
			expectedRes: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GroupAnagrams(tt.inputData)
			if !reflect.DeepEqual(res, tt.expectedRes) {
				t.Errorf("expectedRes: %v, but result: %v", tt.expectedRes, res)
			}
		})
	}
}
