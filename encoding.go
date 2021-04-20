package main

import(
	"fmt"
	"encoding/json"
)

type Message struct{
	Name string
	age int
	address string
}

func main(){
	//str:=`{"name":"bhagya","age":21}`
	m:=Message{"Alice",21,"Andhra pradesh"}
	fmt.Printf("type of struct is :%T\n",m)
	b,err:=json.Marshal(m)
	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println(string(b))
	r := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	fmt.Println(string(r))
	err1 := json.Unmarshal(r, &f)
	if(err1!=nil){
		fmt.Println(err1)
	}
fmt.Println(f)
	// var b
	// fmt.Println(json.Unmarshal(str,&b))
	
	// err:=json.Unmarshal(b,&m)
	// if(err!=nil){
	// 	fmt.Println(err)
	// }
	// fmt.Printf("type of b is :%T\n",b)
	jsondata:=map[string]string{"firstname":"bhagya","lastname":"lakshmi"}
	fmt.Println(jsondata)
	jsonValue,_:=json.Marshal(jsondata)
	fmt.Println(string(jsonValue))
}