package main

import (
	"SofwareGoDay1/data"
	"SofwareGoDay1/humanity"
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

func traiterLesDonnees() {
	mydata, err := data.ReadFile("Go_CSV/medium_csv.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i := range mydata {
		temp, errBis := data.LineToCSV(mydata[i])
		if errBis != nil {
			fmt.Println(err)
		}
		fmt.Println(temp)
	}
}

func créerLaViep1() {
	mydata, err := data.ReadFile("Go_CSV/medium_csv.csv")
	if err != nil {
		fmt.Println(err)
	}
	for i := range mydata {
		temp, errBis := data.LineToCSV(mydata[i])
		if errBis != nil {
			fmt.Println(err)
		}
		human, errBis := humanity.NewHumanFromCSV(temp)
		if errBis != nil {
			fmt.Println(err)
		}
		fmt.Println(human)
	}
}

func créerLaViep2() {
	myHuman, err := humanity.NewHumansFromCsvFile("Go_CSV/medium_csv.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myHuman)
}

func main() {
	// lireUnFichier()
	// traiterLesDonnees()
	// créerLaViep1()
	créerLaViep2()
}
