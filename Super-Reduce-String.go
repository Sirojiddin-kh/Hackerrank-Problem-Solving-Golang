package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "sort"
)

/*
 * Complete the 'superReducedString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func superReducedString(s string) string {
    dict := make(map[string]int)
    var result string
    for _, value := range s {
        dict[string(value)] += 1
    }
    for key, value := range dict {
        if value % 2 != 0 {
            result += key
        }
    }
    result = SortString(result)
    if len(result) != 0 {
        return result
    }
    return "Empty String"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := superReducedString(s)

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
