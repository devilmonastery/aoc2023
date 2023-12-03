package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	colors = []string{"game", "red", "green", "blue"}
)

func read(filepath string) (data []string, err error) {
	f, err := os.Open(filepath)

	if err != nil {
		log.Printf("error reading %q", filepath)
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

type draw struct {
	// index by `colors`
	count []int
}

func (d draw) String() string {
	var buf strings.Builder
	for i := 1; i < len(colors); i++ {
		if i > 1 {
			buf.WriteString(", ")
		}
		buf.WriteString(colors[i])
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%d", d.count[i]))
	}
	buf.WriteString("; ")
	return buf.String()
}

func NewDraw() draw {
	return draw{
		count: make([]int, len(colors)),
	}
}

type game struct {
	number int
	draws  []draw
}

func (g game) String() string {
	return fmt.Sprintf("Game %d: %s", g.number, g.draws)
}

func parseOne(input string) (ret *game, err error) {
	if !strings.HasPrefix(input, "Game ") {
		return nil, fmt.Errorf("invalid line")
	}

	ret = &game{}
	remain := input

	// parse line
	current_num := 0
	current_num_str := ""
	current_draw := NewDraw()

	for _, r := range input {
		//log.Printf("remain: %q", remain)
		// it's a number
		if unicode.IsDigit(r) {
			current_num_str = current_num_str + string(r)
			current_num, err = strconv.Atoi(current_num_str)
			// advance
			remain = remain[1:]
			continue
		}
		//log.Printf("current_num_str:%q current_num: %v", current_num_str, current_num)
		// it's a colon
		if r == ':' {
			//log.Printf("game: %v", current_num)
			ret.number = current_num
			current_num = 0
			current_num_str = ""
		}
		// it's a semicolon
		if r == ';' {
			//log.Printf("new draw")
			ret.draws = append(ret.draws, current_draw)
			current_draw = NewDraw()
		}
		// scan for colors
		color_index := -1
		for i, c := range colors {
			if strings.HasPrefix(remain, c) {
				color_index = i
				break
			}
		}
		if color_index >= 0 {
			//log.Printf("color: %s", colors[color_index])
			current_draw.count[color_index] = current_num
			current_num = 0
			current_num_str = ""
		}
		// advance
		remain = remain[1:]
	}
	// append final draw
	ret.draws = append(ret.draws, current_draw)
	return
}

// Whether a game would be possible for a given number of cubes
func (g *game) Possible(bag ...int) bool {
	if len(bag) != 3 || g == nil {
		return false
	}
	for _, draw := range g.draws {
		for i := 1; i < len(colors); i++ {
			if draw.count[i] > bag[i-1] {
				return false
			}
		}
	}
	return true
}

// Whether a game would be possible for a given number of cubes
func (g *game) Fewest() (ret draw) {
	ret = NewDraw()
	for _, draw := range g.draws {
		for i := 1; i < len(colors); i++ {
			if draw.count[i] > ret.count[i] {
				ret.count[i] = draw.count[i]
			}
		}
	}
	return
}

func (d draw) Power() int {
	return d.count[1] * d.count[2] * d.count[3]
}
