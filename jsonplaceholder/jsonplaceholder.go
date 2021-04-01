package jsonplaceholder

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
)
import "net/http"
import "io/ioutil"


// TODO Write error handler ??
func HandleError(e error)  {

}
type RequestMethod int

const (
	GET = iota
	POST
	PUT
	DELETE
)

func parseRespBody(respBody io.ReadCloser )  (string, error) {
	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		fmt.Println("ERROR:", err)
		return "", err
	}
	return string(body), nil
}

func CreateBody(params url.Values) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	buffer.WriteString(params.Encode())
	return buffer
}

func Query(method RequestMethod , url string, params url.Values, contentType string) (string, error) {
	var resp *http.Response
	var err error

	switch method {
	case GET:
		resp, err = http.Get(url)
	case PUT:
		client := &http.Client{}
		body := CreateBody(params)
		request, err := http.NewRequest("PUT", url, body)
		if err != nil {
			fmt.Println("ERROR:", err)
			return "", err
		}
		resp, err = client.Do(request)
	case POST:
		body := CreateBody(params)
		resp, err = http.Post(url, contentType, body)
	default:
		return "", errors.New("incorrect method")
	}
	if err != nil {
		fmt.Println("ERROR:", err)
		return "", err
	}
	defer resp.Body.Close()
	return parseRespBody(resp.Body)
}