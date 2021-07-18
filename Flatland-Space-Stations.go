package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) int32 {
    
    var (
    m   int32 = int32(len(c))
    i   int32 = 0
    max int32 = 0
)

if n == m {
    return 0
}

for i = 0; i < n; i++ {
    var min int32 = n
    for _, station := range c {
        d := math.Abs(float64(i - station))
        if min > int32(d) {
            min = int32(d)
        }
    }
    if max < min {
        max = min
    }
}

return max


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    cTemp := strings.Split(readLine(reader), " ")

    var c []int32

    for i := 0; i < int(m); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    result := flatlandSpaceStations(n, c)

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
