//quick thing on maps.
//implementing a word count function


package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
// }



 /*
 count the number of distinct words in the input program.
 */
func WordCount(s string) map[string]int {
	var wordsMap map[string]int = make(map[string]int) //MAKE SURE TO INITIALIZE map using make()

	var words []string = strings.Fields(s) //initialize words slice to all words(split by ' ')

	for _,word := range words {
		//notably, if a you try to access a key that does not exist in map, then 
		//go return the zero value of the type of the value in of the map. 
		//in this case, if word is not in wordsMap, then wordsMap[word] returns 0.
		wordsMap[word] = wordsMap[word]+1
	}

	return wordsMap
}


func main(){
	var wordsMap map[string]int = WordCount("hi, ha hah ahaha haha ha ha ha ha ha hahaha my my name is himanshu")

	var total int
	var totalDistinct int
	for k,v := range wordsMap{
		fmt.Println(k,"-->",v)
		total = total + v
		totalDistinct++
	}
	fmt.Println("total words in phrase is", total)
	fmt.Println("total distinct words in phrase is", totalDistinct)
}