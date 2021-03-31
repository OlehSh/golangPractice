package main

import (
	"fmt"
	"github.com/OlehSh/golangPractice/jsonplaceholder"
)

func main()  {
	str, _ := jsonplaceholder.GetPosts()
	fmt.Println(str)
}