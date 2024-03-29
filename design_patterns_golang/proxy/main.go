package main

import "fmt"

/*

----------- DEFINITION: PROXY PATTERN

Proxy is a structural design pattern that provides an object
that acts as a substitute for a real service object used by a
client. A proxy receives client requests, does some work
(access control, caching, etc.) and then passes the request to
a service object.

----------- APPLICATION

- Lazy initialization (virtual proxy). This is when you have a 
heavyweight service object that wastes system resources by being 
always up, even though you only need it from time to time.

- Access control (protection proxy). This is when you want only 
specific clients to be able to use the service object; for instance, 
when your objects are crucial parts of an operating system and 
clients are various launched applications (including malicious ones).

- Logging requests (logging proxy). This is when you want to keep a 
history of requests to the service object.

- Caching request results (caching proxy). This is when you need to 
cache results of client requests and manage the life cycle of this 
cache, especially if results are quite large.

- Smart reference. This is when you need to be able to dismiss a 
heavyweight object once there are no clients that use it.

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
