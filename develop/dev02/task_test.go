package main

import (
	"fmt"
	"testing"
)

func TestUstr(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Тест номер %d", i+1)
		t.Run(testname, func(t *testing.T) {
			ans, err := Ustr(tt.s)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
				t.Error(err)
			}
		})
	}
}
