package day1

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	english2int = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func read(filepath string) (data []string, err error) {
	f, err := os.Open(filepath)

	if err != nil {
		slog.Info("error reading file ", "path", filepath)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		data = append(data, s.Text())
	}

	return
}

func matchEnglish(input string) (val string) {
	for k, v := range english2int {
		if strings.HasPrefix(input, k) {
			return v
		}
	}
	return ""
}

func parseOne(input string) (val int, err error) {
	val = 0
	first := ""
	last := ""
	current := ""
	remain := input
	for _, r := range input {
		// easy case: it's a number
		if unicode.IsDigit(r) {
			current = string(r)
			if first == "" {
				first = current
			}
			last = current
		}
		// scan for an english word
		eng := matchEnglish(remain)
		if eng != "" {
			if first == "" {
				first = eng
			}
			last = eng
		}
		remain = remain[1:]
	}
	if first == "" || last == "" {
		err = fmt.Errorf("no digits found")
		return
	}
	return strconv.Atoi(first + last)
}

func sumAll(filepath string) (val int, err error) {
	data, err := read(filepath)
	if err != nil {
		return
	}

	var tmp int
	for _, line := range data {
		tmp, err = parseOne(line)
		if err != nil {
			log.Printf("parse error on input %q: %v", line, err)
			return
		}
		val += tmp
	}
	return
}
