package gohttp


import (
	"time"
	"bytes"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"errors"
	"strings"
	"net/http"
	"net"
	
)

const (
	defaultMaxIdleConnections = 5
	defaultRequestTimeout = 50 * time.Millisecond
	defaultConnectionTimeout = 5 * time.Second
)


func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	
	//client := http.Client{}

	fullHeaders := c.getReqestHeaders(headers)
	reqestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil{
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	
	request.Header = fullHeaders
	
	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return nil, err
	}

	finalResponse := &Response{
		status : response.Status,
		statusCode : response.StatusCode,
		headers : response.Header,
		body: responseBody,
	}

	return finalResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client{

	c.clientOnce.Do(func() {
		c.client  = &http.Client{
			Timeout : c.getRequestTimeout() + c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost : c.getMaxIdleConnections(),
				ResponseHeaderTimeout : c.getRequestTimeout(),
				DialContext :  (&net.Dialer{
					Timeout : c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}


func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnions > 0 {
		return c.builder.maxIdleConnions
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getRequestTimeout() time.Duration {
	if c.builder.requestTimeout > 0 {
		return c.builder.requestTimeout
	}
	if c.builder.disableTimeouts{
		return 0
	}
	return defaultRequestTimeout
}

func (c * httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts{
		return 0
	}

	return defaultConnectionTimeout
}


func (c *httpClient) getReqestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)
	// Add common headers to the request
	for header, value := range c.builder.headers{
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}    

	// Add custom headers to the request
	for header, value := range requestHeaders{
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}


func (c * httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error){
	if body == nil{
		return nil, nil 
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
		
	case "applicaion/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}


}