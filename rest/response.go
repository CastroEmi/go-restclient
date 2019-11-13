package rest

import (
	"container/list"
	"net/http"
	"sync/atomic"
	"time"
)

// Response structure
type Response struct {
	*http.Response
	Err          error
	byteBody     []byte
	listElement  *list.Element
	ttl          *time.Time
	lastModified *time.Time
	etag         string
	revalidate   bool
	cacheHit     atomic.Value
}

// String return the Response body as a string
func (r *Response) String() string {
	return string(r.Bytes())
}

// Bytes return the Response body as a slice of bytes.
func (r *Response) Bytes() []byte {
	return r.byteBody
}
