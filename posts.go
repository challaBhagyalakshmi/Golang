package main

import (
    "net/http"
    "log"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Posts struct {
    Post []struct{
        UserId int `json:"userId"`
        ID int `json:"id"`
        Title string `json:"title"`
        Body string `json:"body"`
    }
}

func main (){
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

    if err != nil {
        log.Fatal(err)
    }

    content, _ := ioutil.ReadAll(resp.Body)

    var posts Posts

    parsed := json.Unmarshal([]byte(content), &posts)

    //fmt.Println(string(content))

    fmt.Println(parsed)

}