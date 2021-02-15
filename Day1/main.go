package main

import (
	"SofwareGoDay1/data"
	"fmt"
)

func lireUnFichier() {
	mydata, err := data.ReadFile("Go_CSV/medium_csv.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i := range mydata {
		fmt.Println(mydata[i])
	}
}

func traiterlesdonnees() {
	mydata, err := data.LineToCSV("Go_CSV/medium_csv.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i := range mydata {
		fmt.Println(mydata[i])
	}
}

func main() {
	// lireUnFichier()
	traiterlesdonnees()
}
