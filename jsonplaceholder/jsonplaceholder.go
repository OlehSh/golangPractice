package jsonplaceholder

import "fmt"
import "net/http"
import "io/ioutil"

func HandleError(e error)  {

}

func GetPosts() (string, error) {
	url := "https://jsonplaceholder.typicode.com/posts/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println("ERROR:", error)
		return "", error
	}
	str := string(body)
	defer fmt.Println(str)
	return str, nil
}