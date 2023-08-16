// nameAgeSalary.go creates a csv file in the same folder as the program and then writes, reads
// and prints the sorted data to the screen

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func OpenFile(filename string) (*os.File, error) {
	// code below was to test that the function was being called and the file was opening
	// fmt.Println("Opening", filename)
	return os.Open(filename)
}

func CloseFile(f *os.File) {
	// code below was to test that function was being called and the file was closing
	// fmt.Println("Closing file")
	f.Close()
}

func readData(filename string) ([][]string, error) {
	f, err := OpenFile(filename)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	// skip the first row in file which is the column headers
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	// after skipping first row of column headers, reads the rest of the rows in the csv
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	// sorts names by alphabetical order based on the first name
	// if there are two first names that are the same, the program will sort by last name
	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})

	return records, nil
}

func main() {
	defer reportPanic()

	// creates the csv file in the same folder as the program
	f, err := os.Create("employees.csv")
	if err != nil {
		panic(err)
	}
	defer CloseFile(f)

	w := csv.NewWriter(f)

	// employee data
	e1 := Employee{
		Name:   "Tobias Miller",
		Age:    33,
		Salary: 55492.34,
	}

	e2 := Employee{
		Name:   "Helen Bayer",
		Age:    47,
		Salary: 75432.12,
	}

	e3 := Employee{
		Name:   "Abby Recker",
		Age:    22,
		Salary: 33761.85,
	}

	e4 := Employee{
		Name:   "Mike Smith",
		Age:    63,
		Salary: 129106.41,
	}

	e5 := Employee{
		Name:   "Mike Adams",
		Age:    52,
		Salary: 87430.27,
	}

	e6 := Employee{
		Name:   "Bethany Thompson",
		Age:    29,
		Salary: 62497.36,
	}

	// saves employees into an array
	employees := []Employee{e1, e2, e3, e4, e5, e6}

	// creates the column headers in the first row in the csv file
	row := []string{"Name", "Age", "Salary"}
	if err := w.Write(row); err != nil {
		panic(err)
	}

	// writes the employees information to the csv file
	for _, e := range employees {
		row := []string{e.Name, fmt.Sprintf("%d", e.Age), fmt.Sprintf("%.2f", e.Salary)}
		if err := w.Write(row); err != nil {
			panic(err)
		}
	}

	// flushes all data from memory to the file before closing
	// ensures that the write is complete and no corruption at the file level
	w.Flush()

	records, _ := readData("employees.csv")

	// loops through the employee csv and prints the data to the screen
	for _, record := range records {
		fmt.Println("-----------------------------")
		fmt.Println("Name:", record[0])
		fmt.Println("Age:", record[1])
		fmt.Println("Salary:", record[2])
	}
	fmt.Println("-----------------------------")
}
