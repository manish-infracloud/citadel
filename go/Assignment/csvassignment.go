package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type build struct {
	UID, UserID, BRtime, BEtime, BFtime, BIndicat, Bstatus, Bimgsize string
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func repeatCount(list []string) map[string]int {

	RepeatFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the RepeatFrequency map

		_, exist := RepeatFrequency[item]

		if exist {
			RepeatFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			RepeatFrequency[item] = 1 // else start counting from 1
		}
	}
	return RepeatFrequency
}

func getHome() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}

func main() {

	start, _ := time.Parse(time.RFC3339, "2018-10-30T02:47:31-04:00")
	end, _ := time.Parse(time.RFC3339, "2018-11-02T04:06:03-04:00")
	count := 0
	count2 := 0
	var tcount []string
	var UsrID []string

	//Read from CSV file
	filePath := filepath.Join(getHome(), "Downloads", "data.csv")
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//	dateStamp, err := time.Parse(time.RFC3339, date)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var data = build{
			UID:      record[0],
			UserID:   record[1],
			BRtime:   record[2],
			BEtime:   record[3],
			BFtime:   record[4],
			BIndicat: record[5],
			Bstatus:  record[6],
			Bimgsize: record[7],
		}

		if data.Bstatus == "0" {
			count2++
		}
		tcount = append(tcount, data.Bstatus)

		in, _ := time.Parse(time.RFC3339, data.BFtime)
		if inTimeSpan(start, end, in) {
			count++
			//	fmt.Println(in, "is between", start, "and", end, ".")
		}
		UsrID = append(UsrID, data.UserID)
	}

	fmt.Printf("Nubmer of builds beetween the range date from %v to %v are %v\n", start, end, count)

	RepeatMap := repeatCount(UsrID)

	//	for k, v := range RepeatMap {
	//		fmt.Printf("UserID : %s , Count : %d\n", k, v)
	//	}
	//Turning the map into this structure
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range RepeatMap {
		ss = append(ss, kv{k, v})

	}
	//Sorting the key values from repeat map in descending order
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss[:5] {
		fmt.Printf("UserId : %s, Count : %d\n", kv.Key, kv.Value)

	}

	//fmt.Println("Number of successful builds are ", count2)
	//fmt.Println("Total count is:", len(tcount))
	sRate := (float64(count2) / float64(len(tcount))) * 100
	fmt.Printf("Success Rate is: %v%%\n", sRate)
}
