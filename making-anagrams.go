package main 

import (

	"fmt"
	"math"
)

func makingAnagrams(s1 string, s2 string) int32 {
   
   dict := make(map[rune]int)
   var deletions int32

   for _, value := range s1 {

         dict[value] += 1 
   }

   for _, value := range s2 {

           dict[value] -= 1
   }


   for _, value := range dict {

       if value != 0 {
           deletions += int32(math.Abs(float64(value)))
       }

   }


   return deletions
}

func main() {

	s1 := "absdjkvuahdakejfnfauhdsaavasdlkj"
	s2 := "djfladfhiawasdkjvalskufhafablsdkashlahdfa"

	fmt.Println(makingAnagrams(s1, s2))
}
