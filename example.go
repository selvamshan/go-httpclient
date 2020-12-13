package main

import (
	//"time"
	"fmt"
	"github.com/selvamshan/go-httpclient/gohttp"
	//"io/ioutil"
	"net/http"
)

var (
	githubHttpClient = getGithubClient()
)

// User ...
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getGithubClient() gohttp.Client {
	//client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client := gohttp.NewBuilder().
		DisableTimeout(true).
		SetMaxIdleConnections(5).
		SetHeaders(commonHeaders).
		Build()

	//commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer ABC-123")

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

	
}

func createUser(user User) {

	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode())

	//bytes, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}
