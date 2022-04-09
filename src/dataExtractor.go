package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

func generatePieItems(data []string) []opts.PieData {
	hashMap := make(map[string]int)
	for _, ele := range data {
		_, isThere := hashMap[ele]
		if isThere {
			hashMap[ele]++
		} else {
			hashMap[ele] = 1
		}
	}
	items := make([]opts.PieData, 0)

	for str, val := range hashMap {
		fmt.Printf("%v -> %v\n", str, val)
		items = append(items, opts.PieData{
			Name:  str,
			Value: val,
		})
	}

	return items
}

/*
 * @def ageUtility - to extract data regarding age attribute
 * @return error if it exsist
 */
func AgeUtility(df dataframe.DataFrame) {
	datasetAge := df.Col("age")
	ageGroup15_25 := 0
	ageGroup26_35 := 0
	ageGroup36_55 := 0
	ageGroup56_100 := 0
	for i := 0; i < datasetAge.Len(); i++ {
		ele, err := datasetAge.Elem(i).Int()
		if err != nil {
			log.Panicln("Data type Convertion to int")
		}
		if ele > 55 {
			ageGroup56_100++
		} else if ele > 36 {
			ageGroup36_55++
		} else if ele > 26 {
			ageGroup26_35++
		} else if ele > 15 {
			ageGroup15_25++
		}
	}
	// fmt.Printf("%+v\n", datasetAge)
	fmt.Println("Oldest is: ", datasetAge.Max())
	fmt.Println("Youngest is: ", datasetAge.Min())
	fmt.Println("Median is: ", datasetAge.Median())
	fmt.Println("Number People Age 15-25: ", ageGroup15_25)
	fmt.Println("Number People Age 26-35: ", ageGroup26_35)
	fmt.Println("Number People Age 36-55: ", ageGroup36_55)
	fmt.Println("Number People Age 55-100: ", ageGroup56_100)
}

func EmployeeUtility(df dataframe.DataFrame) {

	datasetEmp := df.Col("workclass")
	var probVal float64
	totalNoOfValidRecords := datasetEmp.Len()
	NoOfUnpaidRecords := 0
	NoOfNeverWorkRecords := 0
	NoOfPrivateRecords := 0
	NoOfgovRecords := 0
	arrEmp := make([]string, 0)

	for i := 0; i < datasetEmp.Len(); i++ {
		str := datasetEmp.Elem(i).String()
		arrEmp = append(arrEmp, str)

		switch str {
		case "Private":
			NoOfPrivateRecords++
		// case "Self-emp-not-inc", "Self-emp-inc", "Federal-gov", "Local-gov":
		case "State-gov":
			NoOfgovRecords++
		case "Without-pay":
			NoOfUnpaidRecords++
		case "Never-worked":
			NoOfNeverWorkRecords++
		case "?":
			totalNoOfValidRecords--
		}
	}
	probVal = float64(NoOfUnpaidRecords) / float64(totalNoOfValidRecords)
	fmt.Printf("Percentage of Unpaid workers: %0.5f %%\n", probVal)

	probVal = float64(NoOfNeverWorkRecords) / float64(totalNoOfValidRecords)
	fmt.Printf("Percentage of people who never-worked: %0.5f %%\n", probVal)

	probVal = float64(NoOfPrivateRecords) / float64(totalNoOfValidRecords)
	fmt.Printf("Percentage of private workers: %0.5f %%\n", probVal)

	probVal = float64(NoOfgovRecords) / float64(totalNoOfValidRecords)
	fmt.Printf("Percentage of govnerment workers: %0.5f %%\n", probVal)
	ret := generatePieItems(arrEmp)

	// create a new pie instance
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Employee Work Status",
				Subtitle: "",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("work stat", ret).SetSeriesOptions(
		charts.WithPieChartOpts(
			opts.PieChart{
				Radius: 200,
			},
		),
		charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			},
		),
	)
	f, _ := os.Create("pie.html")
	_ = pie.Render(f)
}
