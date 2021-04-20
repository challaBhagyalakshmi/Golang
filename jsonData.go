package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var err error

	// request http api
	res, err := http.Get("https://api.github.com/repos/challaBhagyalakshmi/Todoapp/commits")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)
	

	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("Body: %s\n", body)

	// parse json
	type jsonUser struct {
		Sha string `json:"sha"`
	}
	user := new(jsonUser)

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello world")
	log.Printf("Received user %s with blog %s", user)
	fmt.Println(user)
	
}
