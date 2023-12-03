package day2

import (
	"log"
	"testing"
)

func Test_read(t *testing.T) {

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
		wantErr bool
		wantNum int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 1},
		{"Game 10: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 10},
		{"Game 99: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 99},
		{"Game 100: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", false, 100},
	}

	for _, tt := range tests {
		val, err := parseOne(tt.input)
		if err == nil && tt.wantErr {
			t.Errorf("input %q, expected error", tt.input)
		}
		if tt.wantNum != val.number {
			t.Errorf("input %q, expected:%v, got:%v", tt.input, tt.wantNum, val.number)
		}
		log.Printf("input:%q val:%s err:%v", tt.input, val, err)
	}
}

func Test_overallPart1(t *testing.T) {
	data, err := read("part1_example.txt")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	total := 0

	for _, line := range data {
		log.Printf("%q", line)
		g, err := parseOne(line)
		if err != nil {
			t.Fatalf("parse error on input %q: %v", line, err)
			return
		}

		ok := g.Possible(12, 13, 14)
		if ok {
			total += g.number
			log.Printf("** OK %d: total:%d", g.number, total)
		} else {
			log.Printf("      %d: total:%d", g.number, total)
		}
	}

	log.Printf("total of games that are ok: %d", total)

	if total != 8 {
		t.Errorf("expected 8, got: %d", total)
	}
}

func Test_answerPart1(t *testing.T) {
	data, err := read("challenge.txt")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	total := 0

	for _, line := range data {
		log.Printf("%q", line)
		g, err := parseOne(line)
		if err != nil {
			t.Fatalf("parse error on input %q: %v", line, err)
			return
		}

		ok := g.Possible(12, 13, 14)
		if ok {
			total = total + g.number
			log.Printf("** OK %d: total:%d", g.number, total)
		} else {
			log.Printf("      %d: total:%d", g.number, total)
		}
	}

	log.Printf("total of games that are ok: %d", total)
}

func Test_overallPart2(t *testing.T) {
	data, err := read("part1_example.txt")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	tests := []draw{
		{count: []int{0, 4, 2, 6}},
		{count: []int{0, 1, 3, 4}},
		{count: []int{0, 20, 13, 6}},
		{count: []int{0, 14, 3, 15}},
		{count: []int{0, 6, 3, 2}},
	}

	total := 0

	for i, line := range data {
		log.Printf("%q", line)
		g, err := parseOne(line)
		if err != nil {
			t.Fatalf("parse error on input %q: %v", line, err)
			return
		}

		d := g.Fewest()
		total = total + d.Power()
		for j := 1; j < len(colors); j++ {
			if tests[i].count[j] != d.count[j] {
				t.Errorf("error on input: %s, expected %v, got %v", line, tests[i], d.count[i])
			}
		}
	}

	if total != 2286 {
		t.Errorf("expected power of 2286, got: %v", total)
	}
}

func Test_answerPart2(t *testing.T) {
	data, err := read("challenge.txt")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	total := 0

	for _, line := range data {
		log.Printf("%q", line)
		g, err := parseOne(line)
		if err != nil {
			t.Fatalf("parse error on input %q: %v", line, err)
			return
		}

		d := g.Fewest()
		log.Printf("fewest: %s", d)
		total = total + d.Power()
	}

	log.Printf("power: %v", total)

}
