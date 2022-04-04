package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	i := 0
	for i < 10000000 {
		data, err := os.Open("adult.csv")
		if err != nil {
			log.Fatal(err)
		}
		df := dataframe.ReadCSV(data)
		fmt.Printf("Dataset[%v]: %v\n", i, df)
		i++
	}
}
