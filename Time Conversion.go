package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	hour, _ := strconv.ParseInt(s[0:2], 0, 64)
	checkingTimeFormat := s[8:9]
	pTimeFormat := strconv.FormatInt((hour+12), 10)
	var militaryTime string = ""
	if checkingTimeFormat == "P" && hour != 12 {
		militaryTime = pTimeFormat + s[2:8]
	}else if checkingTimeFormat == "A" && hour != 12 {
		militaryTime = s[0:8]
	}else {
		militaryTime = "00" + s[2:8]
	}
	return militaryTime

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
