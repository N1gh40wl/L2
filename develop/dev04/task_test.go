package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindQuantities(t *testing.T) {
	var tests = []struct {
		s    []string
		want map[string][]string
	}{
		{[]string{"пятак", "листок", "тяпка", "пятка", "слиток", "столик"}, map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}},
		{[]string{"тест1", "тест2", "тест3"}, map[string][]string{"тест1": {"тест1"}, "тест2": {"тест2"}, "тест3": {"тест3"}}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Тест номер %d", i+1)
		t.Run(testname, func(t *testing.T) {
			ans := FindQuantities(tt.s)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("\ngot  %s\nwant %s", ans, tt.want)
			}
		})
	}
}
