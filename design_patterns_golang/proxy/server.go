package main

// server interface
type server interface {
	handleRequest(string, string) (int, string)
}

// server implementation
type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	} else {
		n.rateLimiter[url] = n.rateLimiter[url] + 1
	}
	// check the number of attempts done after current access
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	return true
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}
