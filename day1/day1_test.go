package day1

import (
	"log"
	"testing"
)

func Test_parse(t *testing.T) {

	data, err := read("part1_example.txt")
	if err != nil {
		t.Errorf("parse() error: %v", err)
	}

	if len(data) == 0 {
		t.Errorf("expected more than zero data")
	}
}

func Test_parseOne(t *testing.T) {
	tests := []struct {
		input   string
		wantVal int
		wantErr bool
	}{
		{"pqr3stu8vwx", 38, false},
		{"two1nine", 29, false},
		{"eightwothree", 83, false},
		{"cow", 0, true},
	}

	for _, tt := range tests {
		val, err := parseOne(tt.input)
		if err == nil && tt.wantErr {
			t.Errorf("input %q, expected error", tt.input)
		}
		if val != tt.wantVal {
			t.Errorf("input %q, expected:%v, got:%v", tt.input, tt.wantVal, val)
		}
	}
}

func Test_sumAllPart1(t *testing.T) {
	val, err := sumAll("part1_example.txt")
	if err != nil {
		t.Errorf("sumAll() error: %v", err)
	}

	if val != 142 {
		t.Errorf("expected 142, got %v", val)
	}

}

func Test_sumAllPart2(t *testing.T) {
	val, err := sumAll("part2_example.txt")
	if err != nil {
		t.Errorf("sumAll() error: %v", err)
	}

	if val != 281 {
		t.Errorf("expected 281, got %v", val)
	}

}

func Test_aoc2023Day1(t *testing.T) {
	val, err := sumAll("challenge.txt")
	if err != nil {
		t.Errorf("sumAll() error: %v", err)
	}

	log.Printf("sum from file:%v", val)
}
