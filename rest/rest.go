package rest

var defaultBuilder = RequestBuilder{}

// Get handles a GET HTTP verb to an specified URL.
//
// In RESTful, GET is used for "reading" or retrieving a resource.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
//
// Get uses the DefaultBuilder.
func Get(url string) *Response {
	//return defaultBuilder.Get(url)
	return nil
}
