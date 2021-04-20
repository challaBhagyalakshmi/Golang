package main

import "fmt"

func main(){
	s:=[]int{1,2,3,4,5}
	fmt.Println(s[1:3])
	arr:=s[2:]
	fmt.Println(arr)
	fmt.Println(s[:5])
	fmt.Println(s[:0])
}