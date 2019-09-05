package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

type build struct {
	UniqueBuildID        string
	UserID               string
	BuildReqReceived     time.Time
	BuildExecutionStart  time.Time
	BuildExecutionEnd    time.Time
	BuildDeleteIndicator bool
	BuildExitCode        int
	BuildImageSize       int
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func repeatitionMap(list []string) map[string]int {

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
	BuildCount := 0
	ExitCodeCount := 0
	IndexNum := 1
	var TotalCount []int
	var userIDList []string
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
		buildReceivedTime, _ := time.Parse(time.RFC3339, record[2])
		buildStartTime, _ := time.Parse(time.RFC3339, record[3])
		buildEndTime, _ := time.Parse(time.RFC3339, record[4])
		buildDeleteStatus, _ := strconv.ParseBool(record[5])
		buildExitStatus, _ := strconv.Atoi(record[6])
		buildImageSize, _ := strconv.Atoi(record[7])

		data := build{
			UniqueBuildID:        record[0],
			UserID:               record[1],
			BuildReqReceived:     buildReceivedTime,
			BuildExecutionStart:  buildStartTime,
			BuildExecutionEnd:    buildEndTime,
			BuildDeleteIndicator: buildDeleteStatus,
			BuildExitCode:        buildExitStatus,
			BuildImageSize:       buildImageSize,
		}
		// Check for the build exit codes, if 0 increment the counter
		if data.BuildExitCode == 0 {
			ExitCodeCount++
		}
		TotalCount = append(TotalCount, data.BuildExitCode)

		in := data.BuildExecutionEnd
		if inTimeSpan(start, end, in) {
			BuildCount++
			//	fmt.Println(in, "is between", start, "and", end, ".")
		}
		userIDList = append(userIDList, data.UserID)

	}

	fmt.Printf("Nubmer of builds beetween the range date from %v to %v are: %d\n", start, end, BuildCount)

	RepeatMap := repeatitionMap(userIDList)

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
	//Print the top 5 values
	fmt.Println("Top 5 users are-")
	for _, kv := range ss[:5] {
		fmt.Printf("%d. UserId : %s, Count : %d\n", IndexNum, kv.Key, kv.Value)
		IndexNum++

	}

	//fmt.Println("Number of successful builds are ", ExitCodeCount)
	//fmt.Println("Total BuildCount is:", len(TotalCount))
	sRate := (float64(ExitCodeCount) / float64(len(TotalCount))) * 100
	fmt.Printf("Success Rate is: %v%%\n", sRate)
}
