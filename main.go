package main

import (
	"fmt"
	"github.com/OlehSh/golangPractice/jsonplaceholder"
	"strconv"
)

const baseUrl string = "https://jsonplaceholder.typicode.com"

func runQuery(i int)  {
	str, _ := jsonplaceholder.Query(jsonplaceholder.GET, baseUrl+"/posts/"+strconv.Itoa(i), nil, "")
	fmt.Println(str)
}

func main()  {
	for i:=1; i<= 10; i++ {
		go runQuery(i)
	}
	var input string
	fmt.Scanln(&input)
}