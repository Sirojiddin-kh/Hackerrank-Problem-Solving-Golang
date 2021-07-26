package main

import (

	"fmt"
)
func palindrome (s string) bool {
    l := len(s)-1
    for i := 0; i <= (l+1)/2; i++ {
        if s[i] != s[l-i] {
            return false
        }
    }
    return true
}
// Complete the palindromeIndex function below.
func palindromeIndex(s string) int32 {
    l := len(s)-1
    for i := 0; i<=l; i, l = i+1,l-1 {
        if s[i] == s[l] {
            continue
        }
        if palindrome(s[i+1:l+1]) {
            return int32(i)
        }
        if palindrome(s[i:l]) {
            return int32(l)
        }
        return -1    
    } 
    return -1
}

func main() {

	str := "baa"

	fmt.Println(palindromeIndex(str))
}