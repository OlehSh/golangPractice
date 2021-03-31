package main

import (
	"example.com/olehsh/jsonplaceholder"
	"fmt"
)

func main()  {
	str:= jsonplaceholder.GetPosts()
	fmt.Println(str)
}