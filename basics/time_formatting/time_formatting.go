package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339)) // 2022-12-02T21:23:50+05:30

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1) // 2012-11-01 22:08:41 +0000 +0000

	// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006
	p(t.Format("3:04 PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	// 9:31 PM
	// Fri Dec  2 21:31:10 2022
	// 2022-12-02T21:31:10.324074+05:30

	t2, _ := time.Parse("3 04 PM", "8 41 PM")
	p(t2) // 0000-01-01 20:41:00 +0000 UTC

	// 2022-12-02T21:32:56-00:00
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// 2022-12-02T21:34:23-00:00

	ansic := "Mon Jan 2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
	// parsing time "8:41PM" as "Mon Jan 2 15:04:05 2006":
	// cannot parse "8:41PM" as "Mon"
}
