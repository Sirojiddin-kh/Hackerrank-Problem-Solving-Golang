package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

)
func ReverseNumber(number int32) int32 {
	strNumber := strconv.Itoa(int(number))
	reverseStrNumber := ""
	for length := len(strNumber); length > 0; length-- {
		reverseStrNumber += string(strNumber[length-1])
	}
	reverseNum, error := strconv.Atoi(reverseStrNumber)
	if error != nil {
		fmt.Println("Failure to cast String to int")
	}
	return int32(reverseNum)
}

/*
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int32, j int32, k int32) int32 {
	var count int32 = 0
	var beautifulDays []int32
	for ; i <= j; i++ {
		beautifulDays = append(beautifulDays, i)
	}
	fmt.Println(beautifulDays)
	for _, value := range beautifulDays {
		if (value - (ReverseNumber(value))) % k == 0 {
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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
