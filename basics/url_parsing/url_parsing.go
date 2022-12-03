package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// scheme, authentication info, host, port, path,
	// query params, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// s := "https://github.com:8000/user?k=1#current"
	u, err := url.Parse(s)
	// Parse the URL and ensure there are no errors
	if err != nil {
		panic(err)
	}
	fmt.Println("Schema: ", u.Scheme)            // psotgres
	fmt.Println("User and Pass: ", u.User)       // user:pass
	fmt.Println("Username: ", u.User.Username()) // user
	fmt.Println(u.User.Password())               // pass
	fmt.Println("Host: ", u.Host)                // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host, port) // host.com 5432

	fmt.Println(u.Path)     // /path
	fmt.Println(u.Fragment) // f

	fmt.Println(u.RawQuery) // k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)         // map[k:[v]]
	fmt.Println(m["k"][0]) // v
}
