package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	data, err := os.Open("adult.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(data)

	fmt.Println(df)
}
