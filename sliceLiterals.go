package main

import "fmt"

func main(){
	s:=[]struct{
		i int
		j bool
	}{
		{1,true},
		{2,false},
		{3,false},
		{4,true},
	}

	fmt.Println(s)
}