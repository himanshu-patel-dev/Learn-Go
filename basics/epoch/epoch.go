package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	// 2022-12-02 09:06:58.367825357 +0530 IST m=+0.000019525

	// Use time.Now with Unix, UnixMilli or UnixNano to get elapsed
	// time since the Unix epoch in seconds, milliseconds or nanoseconds,
	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())
	// 1669952290
	// 1669952290252
	// 1669952290252624331

	// convert integer seconds or nanoseconds since the epoch
	// into the corresponding time.
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))
	// 2022-12-02 09:10:37 +0530 IST
	// 2022-12-02 09:10:37.731147984 +0530 IST
}
