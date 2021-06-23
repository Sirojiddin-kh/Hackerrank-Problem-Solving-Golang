package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sort"

)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxPurchase []int32
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			if keyboards[i] + drives[j] <= b {
				maxPurchase = append(maxPurchase,  keyboards[i] + drives[j])
			}
		}
	}

	sort.Slice(maxPurchase, func(i, j int) bool { return maxPurchase[i] < maxPurchase[j] })

	if len(maxPurchase) == 0 {
		return -1
	}else {
		return maxPurchase[len(maxPurchase) - 1]
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	bnm := strings.Split(readLine(reader), " ")

	bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
	checkError(err)
	b := int32(bTemp)

	nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	keyboardsTemp := strings.Split(readLine(reader), " ")

	var keyboards []int32

	for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
		keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
		checkError(err)
		keyboardsItem := int32(keyboardsItemTemp)
		keyboards = append(keyboards, keyboardsItem)
	}

	drivesTemp := strings.Split(readLine(reader), " ")

	var drives []int32

	for drivesItr := 0; drivesItr < int(m); drivesItr++ {
		drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
		checkError(err)
		drivesItem := int32(drivesItemTemp)
		drives = append(drives, drivesItem)
	}

	/*
	 * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
	 */

	moneySpent := getMoneySpent(keyboards, drives, b)

	fmt.Fprintf(writer, "%d\n", moneySpent)

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
