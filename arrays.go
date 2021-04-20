package main

import "fmt"

func main(){
	var a[10] int
	fmt.Println(a)

	var str [2] string
	str[0]="hello"
	str[1]="world"
	fmt.Println(str)

	var val=[] int{1,2,3,4,5,6}
	fmt.Println(val)
	fmt.Println(val[1:3])

	var var1=[3] int{1,5,8}
	fmt.Println(var1)

	names:=[4]string{"John",
	"Paul",
	"George",
	"Ringo"}

	c:=names[0:2]
	b:=names[2:4]
	fmt.Println(c,b)
}