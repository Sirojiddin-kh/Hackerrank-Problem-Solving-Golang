package main 


func hackerrankInString(s string) string {
    
   // Write your code here
    searchString := "hackerrank"
    i:=0
    for _,j:=range(searchString) {
        p := strings.Index(s[i:], string(j))
        if (p == -1) {
            return "NO"
        }
        i = i + p + 1
    }
    return "YES"
}