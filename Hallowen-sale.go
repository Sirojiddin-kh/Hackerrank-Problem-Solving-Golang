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
 * Complete the 'howManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func howManyGames(p int32, d int32, m int32, s int32) int32 {
    
   var (
       countGames int32
   )
  for i := 0; int32(i) < 10000; i++ {
      s -= p
      if s >=0 {
          countGames +=1
          p -= d
      }
      if p <= m {
          p = m
      }
  }
       
   return countGames
   

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    pTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    p := int32(pTemp)

    dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    m := int32(mTemp)

    sTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
    checkError(err)
    s := int32(sTemp)

    answer := howManyGames(p, d, m, s)

    fmt.Fprintf(writer, "%d\n", answer)

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