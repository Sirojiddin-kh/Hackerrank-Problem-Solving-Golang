package main

import (
	/"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	s = strings.ToLower(s)
	dict := make(map[string]int)

	for _, value := range s {
		dict[string(value)] += 1
	}

	result := len(dict) - 1

	if result == 26 {

		return "pangram"

	} else {

		return "not pangram"
	}

}
