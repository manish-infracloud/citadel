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

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
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
	var Users []string

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

	dup_map := dup_count(UsrID)

	//	for k, v := range dup_map {
	//		fmt.Printf("UserID : %s , Count : %d\n", k, v)
	//	}
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range dup_map {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		//fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		Users = append(Users, kv.Key)

	}
	l := Users[:5]
	fmt.Printf("Top 5 users are: \n%s\n", l)
	//fmt.Println("Number of successful builds are ", count2)
	//fmt.Println("Total count is:", len(tcount))
	sRate := (float64(count2) / float64(len(tcount))) * 100
	fmt.Printf("Success Rate is: %v%%\n", sRate)
}

