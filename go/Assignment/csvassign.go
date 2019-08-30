package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
)

//const (
//	layoutISO = "2006-01-02"
//	layoutUS  = "January 2, 2006"
//)


func main() {

	filePath := filepath.Join(getHome(), "Downloads", "data.csv")

	ioutil.ReadAll
	// Open the file
	csvfile, err := os.Open("/home/manish/Downloads/data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()





	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()

		// func (record) struct object


		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		date := record[4]
		//	dateStamp, err := time.Parse(time.RFC3339, date)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Execution Finished: %s\n", date)

		//fmt.Println(date)
		//	t, _ := time.Parse(layoutUS, date)
		//		fmt.Println(t)
		//fmt.Println(t.Format(layoutUS))
		//fmt.Println(len(record[4]))
	}

	func printslice(slice []string) {
		fmt.Println("slice = ", slice)
	}
	
	func dupCount(list []string) map[string]int {
	
		duplicateFrequency := make(map[string]int)
	
		for _, item := range list {
			// check if the item/element exist in the duplicateFrequency map
	
			_, exist := duplicateFrequency[item]
	
			if exist {
				duplicateFrequency[item]++ // increase counter by 1 if already in the map
			} else {
				duplicateFrequency[item] = 1 // else start counting from 1
			}
		}
		return duplicateFrequency
	}
	
	func main() {
		duplicate := []string
	
		printslice(duplicate)
	
		dupMap := dupCount(duplicate)
	
		//fmt.Println(dupMap)
	
		for k, v := range dupMap {
			fmt.Printf("Item : %s , Count : %d\n", k, v)
		}
	
	}

}


func getHome() string{
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	// windows
	return os.Getenv("USERPROFILE")
}