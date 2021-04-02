package main

import (
	"fmt"
	"github.com/OlehSh/golangPractice/filemanager"
	"github.com/OlehSh/golangPractice/jsonplaceholder"
	"log"
	"strconv"
)

const baseUrl string = "https://jsonplaceholder.typicode.com"

func runQuery(i int)  {
	str, err := jsonplaceholder.Query(jsonplaceholder.GET, baseUrl+"/posts/"+strconv.Itoa(i), nil, "")
	if err != nil {
		log.Fatal(err)
	}
	filemanager.CreateFile(strconv.Itoa(i), str)
	fmt.Println(str)
}

func main()  {
	for i:=1; i <= 10; i++ {
		go runQuery(i)
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
}