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
	fields    string
	delimiter string
	separated bool
}

type Cut struct {
	args
	files  []string
	result []string
}

func CreateCut(fields, delimiter string, separated bool, files []string) *Cut {
	return &Cut{
		args: args{
			fields:    fields,
			delimiter: delimiter,
			separated: separated,
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

func correctLine(line, delimiter string, columns []int) (string, bool) {
	fmt.Println(columns)
	col := strings.Split(line, delimiter)
	isDel := false
	if len(col) > 1 {
		isDel = true
	}
	result := ""
	flag := false
	if isDel {
		for _, v := range columns {
			if (v - 1) < len(col) {
				if flag {
					result += delimiter + col[v-1]
				} else {
					result += col[v-1]
					flag = true
				}
			}
		}
	} else {
		result += line
	}

	return result, isDel
}

func (c *Cut) Run() error {
	files := make([]string, 0)
	for _, v := range c.files {
		file, err := readFile(v)
		if err != nil {
			return err
		}
		files = append(files, file...)
	}
	columns := make([]int, 0)
	columnsString := strings.Split(c.args.fields, ",")
	for _, v := range columnsString {
		c, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		columns = append(columns, c)
	}
	sort.Ints(columns)
	for _, v := range files {
		l, s := correctLine(v, c.args.delimiter, columns)
		if c.args.separated {
			if s {
				c.result = append(c.result, l)
			}
		} else {
			c.result = append(c.result, l)
		}
	}

	return nil
}

func (c *Cut) Output() error {
	for _, v := range c.result {
		fmt.Println(v)
	}
	return nil
}
