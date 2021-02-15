package data

import (
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	mylist := strings.Split(string(data), "\n")
	return mylist, err
}

func LineToCSV(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	mylist := strings.Split(string(data), ",")
	return mylist, err
}
