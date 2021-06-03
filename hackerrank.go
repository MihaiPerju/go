package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	subgraphs := make([]int64, 0)
	subgraphOfNode := make([]int64, n+1, n+1)

	for i := 0; i < len(cities); i++ {
		firstNode := cities[i][0]
		secondNode := cities[i][1]

		if subgraphOfNode[firstNode] == nil && subgraphOfNode[secondNode] == nil {
			subgraphs = append(subgraps, 1)
			subgraphOfNode[firstNode] = &subgraphs(len(subgraphs) - 1)
		} else {

			fmt.Println()
			if subgraphs[firstNode] == 0 {

			} else if subgraphs[secondNode] == 0 {

			} else {
				// they're both from a different graph and they have to be merged
			}
		}

	}

	return numSubgraphs
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("data.out")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nmC_libC_road := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nmC_libC_road[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nmC_libC_road[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(nmC_libC_road[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(nmC_libC_road[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(readLine(reader), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)
		fmt.Println(result)
		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
