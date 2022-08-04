package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type args struct {
	key     string
	number  bool
	reverse bool
	unique  bool
}

type Sort struct {
	args
	files  []string
	result []string
}

func CreateSort(key string, number bool, reverse bool, unique bool, files []string) *Sort {
	return &Sort{
		args: args{
			key:     key,
			number:  number,
			reverse: reverse,
			unique:  unique,
		},
		files: files,
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

func makeUnique(file []string) []string {
	result := make([]string, 0)
	for _, v := range file {
		if isUnique(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func makeReverse(file []string) []string {
	result := make([]string, len(file), cap(file))
	for i, v := range file {
		result[len(file)-i-1] = v
	}
	return result
}

func (s *Sort) Run() error {
	files := make([]string, 0)

	for _, v := range s.files {
		file, err := readFile(v)
		if err != nil {
			return err
		}
		files = append(files, file...)
	}

	if s.args.unique {
		files = makeUnique(files)
	}

	key, err := strconv.Atoi(s.key)
	if err != nil {
		key = 0
	}

	sort.Slice(files, func(i, j int) bool {
		aVals := strings.Split(files[i], " ")
		bVals := strings.Split(files[j], " ")
		if len(aVals) <= key || len(bVals) <= key {
			return false
		}
		if s.args.number {
			aInt, errA := strconv.Atoi(aVals[key])
			bInt, errB := strconv.Atoi(bVals[key])
			if errA == nil && errB == nil {
				return aInt < bInt
			} else if errA != nil && errB != nil {
				return aVals[key] < bVals[key]
			} else if errA != nil {
				return true
			} else if errB != nil {
				return false
			}
		} else {
			return aVals[key] < bVals[key]
		}
		return false
	})

	if s.args.reverse {
		files = makeReverse(files)
	}
	s.result = files
	return nil
}

func (s *Sort) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(s.result, "\n"))
	return err
}
