package filemanager

import (
	"log"
	"io/ioutil"
)

func CreateFile(id string, data string )  {
	filename := "storage/" + id + ".txt"
	err := ioutil.WriteFile(filename, []byte(data), 0664)

	if err != nil {
		log.Fatal(err)
	}
}