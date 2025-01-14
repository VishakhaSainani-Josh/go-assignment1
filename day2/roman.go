package main

import (
	"fmt"
)

func integerValue(romanLetter byte ) int{
	m:=map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	return m[romanLetter]
	
}

func convertRomanToInteger(roman string) int{
	var result int
	for i:=0;i<len(roman);i++ {
		s1:=integerValue(roman[i]);
		if i+1<len(roman) {
			s2:=integerValue(roman[i+1])

			if s2<=s1{
				result+=s1
			} else {
				result+=s2-s1
				i+=1
			}
		} else {
			result+=s1
		}
	}
	return result
}

func main(){
	
	var romanVal string
	fmt.Scanf("%s",&romanVal)

	result:=convertRomanToInteger(romanVal)

	fmt.Println(result)

}