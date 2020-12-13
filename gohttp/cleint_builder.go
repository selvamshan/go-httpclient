package gohttp


import (
	"time"	
	"net/http"
)

type clientBuilder struct{
	maxIdleConnions int
	connectionTimeout time.Duration
	requestTimeout time.Duration
	disableTimeouts bool
	headers http.Header
	baseUrl string
}

// ClientBuilder ..
type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(time.Duration)  ClientBuilder
	SetRequestTimeout(time.Duration)  ClientBuilder
	SetMaxIdleConnections(int)  ClientBuilder
	DisableTimeout(bool)  ClientBuilder
	Build() Client 
}

// NewBuilder ...
func NewBuilder() ClientBuilder{
	builder := &clientBuilder{}
	return builder
}


func (c *clientBuilder) Build() Client {
	client := &httpClient{
		builder: c,
	}
	return client
}

func (c *clientBuilder) SetHeaders(headers http.Header)  ClientBuilder { 
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetRequestTimeout(timeout time.Duration)  ClientBuilder {
	c.requestTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(i int)  ClientBuilder {
	c.maxIdleConnions = i
	return c
}

func (c *clientBuilder) DisableTimeout(disable bool)  ClientBuilder {
	c.disableTimeouts = disable
	return c
}