package data

import (
	"errors"
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	mylist := strings.Split(string(data), "\n")
	return mylist, err
}

func LineToCSV(line string) ([]string, error) {
	mylist := strings.Split(line, ",")
	if len(mylist) != 3 {
		return nil, errors.New("Bad format")
	}
	return mylist, nil
}

// func Add(a, b fmt.Stringer) string {
// 	return a.String() + b.String()
// }
