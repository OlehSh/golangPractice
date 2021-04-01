package main

import (
	"fmt"
	"github.com/OlehSh/golangPractice/jsonplaceholder"
)

const baseUrl string = "https://jsonplaceholder.typicode.com"

func main()  {
	str, _ := jsonplaceholder.Query( jsonplaceholder.GET, "https://jsonplaceholder.typicode.com/posts", nil,"")
	fmt.Println(str)
}