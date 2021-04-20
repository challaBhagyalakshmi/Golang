package main

import "fmt"

func main(){
	s:=[]int{1,2,3,4,5,6}

	fmt.Printf("length of s is :%d\ncapacity of s is:%d\n %v\n",len(s),cap(s),s)

	p:=[9]string{"bhagya","lakshmi"}
	fmt.Printf("length of p is :%d\ncapacity of s is:%d\n %v\n",len(p),cap(p),p)

	var arr[] int

	if(arr==nil){
		fmt.Println("nil!")
	}

	fmt.Println(arr,len(arr),cap(arr));
}