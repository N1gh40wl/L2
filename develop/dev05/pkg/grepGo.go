package pkg

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type args struct {
	after      string
	before     string
	context    string
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

type Grep struct {
	args
	files   []string
	pattern string
	result  []string
}

func CreateGrep(after, before, context string, count, ignoreCase, invert, fixed, lineNum bool, files []string, pattern string) *Grep {
	return &Grep{
		args: args{
			after:      after,
			before:     before,
			context:    context,
			count:      count,
			ignoreCase: ignoreCase,
			invert:     invert,
			fixed:      fixed,
			lineNum:    lineNum,
		},
		files:   files,
		pattern: pattern,
	}
}

func readFile(filepath string) ([]string, error) {
	result := make([]string, 0)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		result = append(result, string(line))
	}
	return result, nil
}

func isUnique(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return false
		}
	}
	return true
}

func FindPattern(s []string, pattern string, register bool) ([]int, error) {
	result := make([]int, len(s), cap(s))
	err := errors.New("No match")
	sCopy := make([]string, 0)
	if register {
		pattern = strings.ToLower(pattern)
	}
	for i, v := range s {
		if register {
			sCopy = append(sCopy, strings.ToLower(v))
		} else {
			sCopy = append(sCopy, v)
		}
		result[i] = strings.Index(sCopy[i], pattern)
		if result[i] != -1 {
			err = nil
		}
	}
	return result, err
}

func countLines(lines []int) int {
	result := 0
	for _, v := range lines {
		if v != -1 {
			result++
		}
	}
	return result
}

func markCorrectLines(lines []int, up, down int) []bool {
	result := make([]bool, len(lines), cap(lines))
	for i, v := range lines {
		if v != -1 {
			left := i - up
			right := i + down
			if left < 0 {
				left = 0
			}
			if right > len(lines)-1 {
				right = len(lines) - 1
			}
			for j := left; j <= right; j++ {
				result[j] = true
			}
		}
	}
	return result
}

func (g *Grep) Run() error {

	files := make([]string, 0)

	for _, v := range g.files {
		file, err := readFile(v)
		if err != nil {
			return err
		}
		files = append(files, file...)
	}
	//
	lineIndex, err := FindPattern(files, g.pattern, g.args.ignoreCase)
	if err != nil {
		return err
	}

	if g.args.fixed {
		flag := false
		for _, v := range lineIndex {
			if v != -1 {
				if files[v] == g.pattern {
					flag = true
				} else {
					v = -1
				}
			}
		}
		if !flag {
			return errors.New("No match")
		}
	}

	if g.args.count {
		g.result = append(g.result, strconv.Itoa(countLines(lineIndex)))
		return nil
	}

	if g.args.invert {
		for i, v := range files {
			if lineIndex[i] == -1 {
				g.result = append(g.result, v)
			}
		}
		return nil
	}

	after, err := strconv.Atoi(g.after)

	if err != nil {
		after = 0
	}
	before, err := strconv.Atoi(g.before)
	if err != nil {
		before = 0
	}
	contextV, err := strconv.Atoi(g.context)
	if err != nil {
		contextV = 0
	}
	if contextV > after {
		after = contextV
	}
	if contextV > before {
		before = contextV
	}
	finalLines := markCorrectLines(lineIndex, before, after)
	for i, v := range finalLines {
		if v {
			resline := ""
			if g.args.lineNum {
				resline = strconv.Itoa(i+1) + "- " + files[i]
			} else {
				resline = files[i]
			}

			g.result = append(g.result, resline)

		}
	}
	return nil
}

func (g *Grep) Output() error {
	for _, v := range g.result {
		println(v)
	}
	return nil
}
