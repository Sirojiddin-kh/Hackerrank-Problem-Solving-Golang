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
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	inputLength := int64(len(s))
	occurrenceOfInput := occurrenceOfA(s)
	numberOfSubstring := int64(n / inputLength)

	numberOfAs := numberOfSubstring * occurrenceOfInput
	reminder := n % inputLength

	for i := int64(0); i < reminder; i++ {
		if s[i] == 'a'{
			numberOfAs += 1
		}
	}
	return numberOfAs
}

func occurrenceOfA(s string) int64 {
	var count int64
	for i := range s {
		if s[i] == 'a'{
			count += 1
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
