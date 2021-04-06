package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/OlehSh/golangPractice/jsonplaceholder"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

const baseUrl string = "https://jsonplaceholder.typicode.com"

var db *sql.DB

type PostMessage struct {
	UserId, Id  int
	Title, Body string
}

type CommentMessage struct {
	PostId, Id        int
	Name, Email, Body string
}

func parseJson(body string, i interface{}) {
	err := json.Unmarshal([]byte(body), &i)

	if err != nil {
		log.Fatal(err)
	}
}

func runQuery(i int) {

	str, err := jsonplaceholder.Query(jsonplaceholder.GET, baseUrl+"/comments/?postId="+strconv.Itoa(i), nil, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Comment =>", str)
	var respJson []CommentMessage
	parseJson(str, &respJson)
	saveComment(respJson)
}

func savePosts(data PostMessage) {
	query := fmt.Sprintf("INSERT INTO posts (post_id, user_id, body, title) VALUES (%d, %d, '%s', '%s' )",
		data.Id, data.UserId, data.Body, data.Body)
	fmt.Println("QUERY", query)
	fmt.Println("SAVE POST", data.Id)
	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
}
func saveComment(data []CommentMessage) {
	fmt.Println("SAVE COMMENT")
	for _, c := range data {
		query := fmt.Sprintf("INSERT INTO comments (comment_id, post_id, name, email, body) VALUES (%d, %d, '%s', '%s', '%s' )",
			c.Id, c.PostId, c.Name, c.Email, c.Body)
		fmt.Println(query)
		db.Query(query)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "gopractice"
	password = "qwerty"
	dbname   = "gopractice"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("DB connect:", err)
	}
	str, err := jsonplaceholder.Query(jsonplaceholder.GET, baseUrl+"/posts/?userId=7", nil, "")
	if err != nil {
		log.Fatal(err)
	}
	var posts []PostMessage
	parseJson(str, &posts)

	for _, post := range posts {
		go savePosts(post)
		go runQuery(post.Id)
	}
	var input string
	fmt.Scanln(&input)

}
