package rest

import (
	"net/http"
	"sync"
	"time"
)

// The default transport used by all RequestBuilders
// that haven't set up a CustomPool
var defaultTransport http.RoundTripper

// Sync once to set default client and transport to default Request Builder
var dTransportMtxOnce sync.Once

// DefaultTimeout is the default timeout for all clients.
// DefaultConnectTimeout is the time it takes to make a connection
// Type: time.Duration
var DefaultTimeout = 500 * time.Millisecond

// DefaultConnectTimeout ...
var DefaultConnectTimeout = 1500 * time.Millisecond

// DefaultMaxIdleConnsPerHost is the default maxium idle connections to have
// per Host for all clients, that use *any* RequestBuilder that don't set
// a CustomPool
var DefaultMaxIdleConnsPerHost = 2

// ContentType represents the HTTP Vebs body's type
type ContentType int

const (
	// JSON represents a JSON content type
	JSON ContentType = iota
	// XML represents an XML content type
	XML
	// BYTES represents a plain content type
	BYTES
)

// RequestBuilder is the baseline to create requests
// There is a DefaultBuilder that you may use for simple requests
// RequestBuilder is thread-safe and should be stored for later usage.
type RequestBuilder struct {

	// Headers to be send in the request.
	Headers http.Header

	// Timeout to complete request
	Timeout time.Duration

	// ConnectionTimeout bounds the time spent obtaining a successful connection
	ConnectionTimeout time.Duration

	// Base url is used to build the entire url for make the request (BaseURL + URI)
	BaseURL string

	// ContentType
	ContentType ContentType

	// Disable internal caching of response
	DisableCache bool

	// Disable timeout.
	DisableTimeout bool

	// Set the http client to follow a redirect if it gets a 3xx response from the server
	FollowRedirect bool

	// CustomPool create a pool if you don't want to share the *transport*, with others RequestBuilders
	CustomPool *CustomPool

	// Set basic Auth created request
	BasicAuth *BasicAuth

	// Set a specific user agent for the created request
	UserAgent string

	// Public for custom fine tuning
	Client *http.Client

	clientMtxOnce sync.Once
}

// CustomPool defines a separated internal *transport* and connection pooling.
type CustomPool struct {
	MaxIdleConnsPerHost int
	Proxy               string

	// Public for custom fine tuning
	Transport http.RoundTripper
}

// BasicAuth allows to set UserName and Password for a given RequestBuilder
type BasicAuth struct {
	UserName string
	Password string
}

// Get ...
func (rb *RequestBuilder) Get(url string) *Response {
	return rb.doRequest(http.MethodGet, url, nil)
}

// Post ...
func (rb *RequestBuilder) Post(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPost, url, body)
}

// Put ...
func (rb *RequestBuilder) Put(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPut, url, body)
}

// Delete ...
func (rb *RequestBuilder) Delete(url string) *Response {
	return rb.doRequest(http.MethodDelete, url, nil)
}

// Patch ...
func (rb *RequestBuilder) Patch(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPatch, url, body)
}

// Head ...
func (rb *RequestBuilder) Head(url string) *Response {
	return rb.doRequest(http.MethodHead, url, nil)
}

// Options ...
func (rb *RequestBuilder) Options(url string) *Response {
	return rb.doRequest(http.MethodOptions, url, nil)
}

// AsyncGet ...
func (rb *RequestBuilder) AsyncGet(url string, f func(*Response)) {
	go doAsyncRequest(rb.Get(url), f)
}

// AsyncPost ...
func (rb *RequestBuilder) AsyncPost(url string, body interface{}, f func(*Response)) {
	go doAsyncRequest(rb.Post(url, body), f)
}

// AsyncPut ...
func (rb *RequestBuilder) AsyncPut(url string, body interface{}, f func(*Response)) {
	go doAsyncRequest(rb.Put(url, body), f)
}

// AsyncPatch ...
func (rb *RequestBuilder) AsyncPatch(url string, body interface{}, f func(*Response)) {
	go doAsyncRequest(rb.Patch(url, body), f)
}

// AsyncDelete ...
func (rb *RequestBuilder) AsyncDelete(url string, f func(*Response)) {
	go doAsyncRequest(rb.Delete(url), f)
}

// AsyncHead ...
func (rb *RequestBuilder) AsyncHead(url string, f func(*Response)) {
	go doAsyncRequest(rb.Head(url), f)
}

// AsyncOptions ...
func (rb *RequestBuilder) AsyncOptions(url string, f func(*Response)) {
	go doAsyncRequest(rb.Options(url), f)
}

func doAsyncRequest(r *Response, f func(*Response)) {
	f(r)
}

// ForkJoin let you *fork* requests, and *wait* until all of them have return.
//
// 	var futureA, futureB *rest.FutureResponse
//
// 	rest.ForkJoin(func(c *rest.Concurrent){
//		futureA = c.Get("/url/1")
//		futureB = c.Get("/url/2")
//	})
//
//	fmt.Println(futureA.Response())
//	fmt.Println(futureB.Response())
//
func (rb *RequestBuilder) ForkJoin(f func(c *Concurrent)) {

	c := new(Concurrent)
	c.reqBuilder = rb

	f(c)

	c.wg.Add(c.list.Len())

	for e := c.list.Front(); e != nil; e = e.Next() {
		go e.Value.(func())()
	}

	c.wg.Wait()
}
