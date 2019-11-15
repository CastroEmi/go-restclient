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

// Post handles a POST HTTP verb to an specified URL.
//
// In RESTful, POST is used to send user-generated data to the server.
// This data is stored in a subordinate of the resourse identified in the
// Request-URI. In other words, POST append a resourse at the end of a given url.
// Client should expect a response status code of 201(Created), 400(Bad Request),
// 404(Not Found), 405(Method Not Allowed) or 409(Conflict) if resource already exist.
//
// Body could be any of the form: string, []byte, struct & map.
func Post(url string, body interface{}) *Response {
	//return defaultBuilder.Post(url, body)
	return nil
}

// Put handles a PUT HTTP verb to an specified URL
//
// In RESTful, PUT is used to send user-generated data to the server.
// Knowing the exact URI of the resourse, it creates it or, if the resourse
// already exists it ovewrites it with the Request-body data.
// PUT verb is idempotent, it means that if you apply it n times it
// will always return the same result.
// Client should expect a response status code of 200(OK), 201(Created), 204(No Content),
// 400(Bad Request), 404(Not Found), 405(Method Not Allowed).
//
// Body could be any of the form: string, []byte, struct & map.
//
// Put uses the DefaultBuilder.
func Put(url string, body interface{}) *Response {
	//return defaultBuilder.Put(url, body)
	return nil
}

// Patch issues a PATCH HTTP verb to the specified URL
//
// In Restful, PATCH is used for "partially updating" a resource.
// Client should expect a response status code of of 200(OK), 404(Not Found),
// or 400(Bad Request). 200(OK) could be also 204(No Content)
//
// Body could be any of the form: string, []byte, struct & map.
//
// Patch uses the DefaultBuilder.
func Patch(url string, body interface{}) *Response {
	return defaultBuilder.Patch(url, body)
}

// Delete handles a DELETE HTTP verb to an specified URL.
//
// In RESTful, DELETE is used to "delete" an specified resource from the server.
// Client should expect a response status code of 200(OK),
// 400(Bad Request), 401(Unauthorized), 403(Forbiden), 404(Not Found), 405(Method Not Allowed).
//
// Delete uses the DefaultBuilder.
func Delete(url string) *Response {
	//return defaultBuilder.Delete(url)
	return nil
}

// Head issues a HEAD HTTP verb to the specified URL
//
// In Restful, HEAD is used to "read" a resource headers only.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
//
// Head uses the DefaultBuilder.
func Head(url string) *Response {
	return defaultBuilder.Head(url)
}

// Options issues a OPTIONS HTTP verb to the specified URL
//
// In Restful, OPTIONS is used to get information about the resource
// and supported HTTP verbs.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
func Options(url string) *Response {
	return defaultBuilder.Options(url)
}

// AsyncGet is the *asynchronous* option for GET.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncGet uses the DefaultBuilder
func AsyncGet(url string, f func(*Response)) {
	defaultBuilder.AsyncGet(url, f)
}

// AsyncPost is the *asynchronous* option for POST.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncPost uses the DefaultBuilder
func AsyncPost(url string, body interface{}, f func(*Response)) {
	defaultBuilder.AsyncPost(url, body, f)
}

// AsyncPut is the *asynchronous* option for PUT.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncPut uses the DefaultBuilder
func AsyncPut(url string, body interface{}, f func(*Response)) {
	defaultBuilder.AsyncPut(url, body, f)
}

// AsyncDelete is the *asynchronous* option for DELETE.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncDelete uses the DefaultBuilder
func AsyncDelete(url string, f func(*Response)) {
	defaultBuilder.AsyncDelete(url, f)
}

// AsyncPatch is the *asynchronous* option for PATCH.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncPatch uses the DefaultBuilder
func AsyncPatch(url string, body interface{}, f func(*Response)) {
	defaultBuilder.AsyncPatch(url, body, f)
}

// AsyncHead is the *asynchronous* option for HEAD.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncHead uses the DefaultBuilder
func AsyncHead(url string, f func(*Response)) {
	defaultBuilder.AsyncHead(url, f)
}

// AsyncOptions is the *asynchronous* option for OPTIONS.
// The go routine calling AsyncGet(), will not be blocked.
//
// Whenever the Response is ready, the *f* function will be called back.
//
// AsyncOptions uses the DefaultBuilder
func AsyncOptions(url string, f func(*Response)) {
	defaultBuilder.AsyncOptions(url, f)
}

// ForkJoin let you *fork* requests, and *wait* until all of them have return.
//
func ForkJoin(f func(*Concurrent)) {
	defaultBuilder.ForkJoin(f)
}
