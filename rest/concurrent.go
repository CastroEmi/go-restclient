package rest

import (
	"container/list"
	"net/http"
	"sync"
	"sync/atomic"
	"unsafe"
)

// FutureResponse represents a response to be completed after a ForkJoin operation
// is done
//
// FutureResponse will never be nil and has a Response fucntion for getting the Response
// that will be nil after the ForkJoin operation is completed.
type FutureResponse struct {
	p unsafe.Pointer
}

// Response gives you the Response of a Request after the ForkJoin operation is completed.
// Response will be nil if the ForkJoin operation is not completed.
func (fr *FutureResponse) Response() *Response {
	return (*Response)(fr.p)
}

// Concurrent implements methods for HTTP verbs.
//The difference with synchronous API is that concurrent methods return a FutureResponse
// which hold a pointer to a Response, which is nil until the request has finished.
type Concurrent struct {
	list       list.List
	wg         sync.WaitGroup
	reqBuilder *RequestBuilder
}

// Get perform a GET HTTP verb to the specified URL concurrently.
func (c *Concurrent) Get(url string) *FutureResponse {
	return c.doRequest(http.MethodGet, url, nil)
}

// Post perform a POST HTTP verb to the specified URL concurrently.
//
// Body could be any of the form: string, []byte, struct & map.
func (c *Concurrent) Post(url string, body interface{}) *FutureResponse {
	return c.doRequest(http.MethodPost, url, body)
}

// Put perform a PUT HTTP verb to the specified URL concurrently.
//
// Body could be any of the form: string, []byte, struct & map.
func (c *Concurrent) Put(url string, body interface{}) *FutureResponse {
	return c.doRequest(http.MethodPut, url, body)
}

// Patch perform a PATCH HTTP verb to the specified URL concurrently.
//
// Body could be any of the form: string, []byte, struct & map.
func (c *Concurrent) Patch(url string, body interface{}) *FutureResponse {
	return c.doRequest(http.MethodPatch, url, body)
}

// Delete perform a DELETE HTTP verb to the specified URL concurrently.
func (c *Concurrent) Delete(url string) *FutureResponse {
	return c.doRequest(http.MethodDelete, url, nil)
}

// Head perform a HEAD HTTP verb to the specified URL concurrently.
func (c *Concurrent) Head(url string) *FutureResponse {
	return c.doRequest(http.MethodHead, url, nil)
}

// Options perform a OPTIONS HTTP verb to the specified URL concurrently.
func (c *Concurrent) Options(url string) *FutureResponse {
	return c.doRequest(http.MethodOptions, url, nil)
}

func (c *Concurrent) doRequest(verb string, url string, body interface{}) *FutureResponse {
	fr := new(FutureResponse)

	future := func() {
		defer c.wg.Done()
		response := c.reqBuilder.doRequest(verb, url, body)
		atomic.StorePointer(&fr.p, unsafe.Pointer(response))
	}

	c.list.PushBack(future)

	return fr
}
