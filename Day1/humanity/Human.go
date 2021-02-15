package humanity

import (
	"SofwareGoDay1/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	if len(csv) != 3 {
		return nil, errors.New("Bad format")
	}
	tempAge, err := strconv.Atoi(csv[1])
	if err != nil {
		return nil, errors.New("Not int")
	}
	human := Human{csv[0], tempAge, csv[2]}
	return &human, err
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	fileContent, err := data.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lenContent := len(fileContent)
	human := make([]*Human, lenContent)
	for i, line := range fileContent {
		lineCsv, err := data.LineToCSV(line)
		if err != nil {
			return nil, err
		}
		human[i], err = NewHumanFromCSV(lineCsv)
		if err != nil {
			return nil, err
		}
	}
	return human, err
}

func NewHumansFromJsonFile(path string) ([]*Human, error) {
	var human []*Human
	temp, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("Error read")
	}
	bytes := []byte(temp)
	err = json.Unmarshal(bytes, &human)
	if err != nil {
		return nil, errors.New("Error Json")
	}
	return human, err
}

func (h *Human) String() string {
	return fmt.Sprintf("Name : %s, Age %d, Country : %s\n", h.Name, h.Age, h.Country)
}
