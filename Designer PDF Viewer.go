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
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

func designerPdfViewer(h []int32, word string) int32 {
	arr := strings.Split(word, "")
	letters := map[string]int32 {
		"a" : h[0],
		"b" : h[1],
		"c" : h[2],
		"d" : h[3],
		"e" : h[4],
		"f" : h[5],
		"g" : h[6],
		"h" : h[7],
		"i" : h[8],
		"j" : h[9],
		"k" : h[10],
		"l" : h[11],
		"m" : h[12],
		"n" : h[13],
		"o" : h[14],
		"p" : h[15],
		"q" : h[16],
		"r" : h[17],
		"s" : h[18],
		"t" : h[19],
		"u" : h[20],
		"v" : h[21],
		"w" : h[22],
		"x" : h[23],
		"y" : h[24],
		"z" : h[25],
	}
	var words []int32
	for i := 0; i < len(arr)  -1; i++ {
		for key, _:= range letters {
			if arr[i] == key {
				words = append(words, letters[key])

			}
		}
	}
	max := words[0]
	for _, value := range words {
		if max < value {
			max = value
		}
	}
	return max * int32(len(arr))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var h []int32

	for i := 0; i < 26; i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	word := readLine(reader)

	result := designerPdfViewer(h, word)

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
