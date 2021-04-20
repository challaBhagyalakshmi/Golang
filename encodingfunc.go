package main

import (
    "net/http"
    "log"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Post struct {
        UserId int `json:"userId"`
        ID int `json:"id"`
        Title string `json:"title"`
        Body string `json:"body"`
}

type Posts []Post


func main (){
    resp, err := http.Get("https://api.github.com/repos/challaBhagyalakshmi/Todoapp/commits")

    if err != nil {
        log.Fatal(err)
    }

    content, _ := ioutil.ReadAll(resp.Body)

    var posts Posts

    err = json.Unmarshal(content, &posts)

    if err != nil {
        log.Fatal(err)
    }


    fmt.Println(string(posts))

}