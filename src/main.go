package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pie.html")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the Server started 🚀")
	fmt.Fprintf(w, "Try using /pie in url!!")
}

func main() {

	// 0 for 1
	for i := 0; i < 10000000; i++ {
		data, err := os.Open("adult.csv")
		if err != nil {
			log.Fatal(err)
		}
		df := dataframe.ReadCSV(data)
		fmt.Printf("Dataset[%v]: %v\n", i, df)
	}
	data, err := os.Open("adult.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(data)
	fmt.Println("Attributes: ", df.Names())
	AgeUtility(df)
	EmployeeUtility(df)
	http.HandleFunc("/pie", html)
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":8080", nil)
}
