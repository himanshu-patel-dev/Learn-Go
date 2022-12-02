package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	// 2022-12-02 08:48:15.985332417 +0530 IST m=+0.000029683

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 2009-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	// 2009
	// November
	// 17
	// 20
	// 34
	// 58
	// 651387237
	// UTC
	// Tuesday

	p(now.Location()) // Local
	p(now.Zone())     // IST 19800

	p(then.Before(now)) // true
	p(then.Equal(now))  // false
	p(then.After(now))  // false

	diff := now.Sub(then)
	p(diff) // 114294h49m50.335947221s

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// 114294.83666685014
	// 6.857690200011008e+06
	// 4.1146141200066054e+08
	// 411461412000660523

	p(then.Add(diff))  // 2022-12-02 03:26:06.567153435 +0000 UTC
	p(then.Add(-diff)) // 1996-11-03 13:43:50.735621039 +0000 UTC
	// why don't we use then.Sub to subtract the time from then
	// because Sub accept time.Time to sub one time from another
	// while  diff is time.Duration

}
