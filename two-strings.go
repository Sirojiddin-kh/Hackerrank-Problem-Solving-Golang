package main


func twoStrings(s1 string, s2 string) string {
    
    dict := make(map[string]bool)
    var substring string
    for _, value := range s1 {
        
        dict[string(value)] = true
    }
    
    for _, value := range s2 {
        
        if dict[string(value)] {
            
            substring += string(value)
        }
    }
    
    if len(substring) > 0 {
        return "YES"
    } else {
        return "NO"
    }

}
