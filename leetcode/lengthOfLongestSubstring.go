package leetcode


func charInMap(char rune, charMap map[rune]bool ) bool {
	_, ok := charMap[char];
	return ok
}

func LengthOfLongestSubstring(str string) int{
	if str == ""{
		return 0
	}
	max := 1
	for i,val := range str{
		chars := map[rune]bool{}
		chars[val] = true
		for j := i+1; j < len(str); j++ {
			if(charInMap(rune(str[j]), chars)){
				break
			} else {
				chars[rune(str[j])] = true
				length := j-i +1
				if length > max {
					max = length
				} 
			}
			
			
		}
		
	}
	return max

}
