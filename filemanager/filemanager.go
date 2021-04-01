package filemanager

import (
	"log"
	"os"
)

func CreateFile(id string, data string )  {
	file, err := os.Create("storage/" + id + ".txt")

	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(data)

	if err != nil {
		log.Fatal(err)
	}
}