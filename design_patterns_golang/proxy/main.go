package main

import "fmt"

/*

Proxy is a structural design pattern that provides an object
that acts as a substitute for a real service object used by a
client. A proxy receives client requests, does some work
(access control, caching, etc.) and then passes the request to
a service object.

*/

func main() {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	// these call can go to applicatio directly and call the handleRequest on
	// application itself, but calling same method on proxy ensure
	// ratelimit is respected by clients
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	// exhausted limit of 2 request
	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}

/*

Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 200
Body: Ok

Url: /app/status
HttpCode: 403
Body: Not Allowed

Url: /app/status
HttpCode: 201
Body: User Created

Url: /app/status
HttpCode: 404
Body: Not Ok

*/
