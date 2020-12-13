package gohttp

import (
	"net/http"
	"testing"
	//"fmt"
)


func TestRequestHeaders(t  *testing.T) {
	client:= httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getReqestHeaders(requestHeaders)

	if len(finalHeaders) != 3{
		t.Error("exptect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content Type Recieved")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid User Agent Recieved")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid User Id Recieved")
	}

}

func TestRequestBody(t *testing.T){
	// Intialize
	client:= httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T){
		// Execute
		body, err := client.getRequestBody("", nil)
		//Validate
		if err != nil{
			t.Error("No error exptected when passing nil body")
		}

		if body != nil{
			t.Error("No body expected when passing nil body")
		}
	})
	
	t.Run("BodyWithJson", func(t *testing.T){
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		// fmt.Println(err)
		// fmt.Println(string(body))
		if err != nil {
			t.Error("No error exptected when marsheling slice as json body")
		}

		if string(body) != `["one","two"]` {
			t.Error("Invalid json obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T){})

	t.Run("DefaultJsonBody", func(t *testing.T){})

	
	
}
