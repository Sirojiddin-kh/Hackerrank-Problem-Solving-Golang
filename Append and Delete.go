package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	m := make(map[string]bool)
	var sString, tString []string
	sString = strings.Fields(s)
	tString = strings.Fields(t)
	var similarStrings []string

	for _, item := range sString {
		m[item] = true
	}

	for _, item := range tString {
		if _, ok := m[item]; ok {
			similarStrings = append(similarStrings, item)
		}
	}
	var appended int
	delete := len(sString) - len(similarStrings)
	if len(tString) == len(similarStrings) {
		appended = len(similarStrings)
	}else{
		appended = len(tString) - len(similarStrings)
	}

	if int32(delete) + int32(appended) < k {
		return "Yes"
	}else {
		return "No"
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

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