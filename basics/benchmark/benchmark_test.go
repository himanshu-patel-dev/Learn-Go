// 1. To test a file that should be part of some module
// go mod init benchmark && go mod tidy

// 2. File name should end with _test.go suffix for the tests
// to be picked up.

// For the sake of demonstration, this code is in package main,
// but it could be any package. Testing code typically lives in
// the same package as the code it tests
package main

import (
	"fmt"
	"testing"
)

// Typically, the code we’re testing would be in a source file
// named something like intutils.go, and the test file for it
// would then be named intutils_test.go
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name
// beginning with Test.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* will report test failures but continue
		// executing the test. t.Fatal* will report te
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run enables running “subtests”, one for each table entry.
		// These are shown separately when executing go test -v.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// $ go test -v
// === RUN   TestIntMinBasic
// --- PASS: TestIntMinBasic (0.00s)
// === RUN   TestIntMinTableDriven
// === RUN   TestIntMinTableDriven/0,1
// === RUN   TestIntMinTableDriven/1,0
// === RUN   TestIntMinTableDriven/2,-2
// === RUN   TestIntMinTableDriven/0,-1
// === RUN   TestIntMinTableDriven/-1,0
// --- PASS: TestIntMinTableDriven (0.00s)
//     --- PASS: TestIntMinTableDriven/0,1 (0.00s)
//     --- PASS: TestIntMinTableDriven/1,0 (0.00s)
//     --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
//     --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
//     --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
// PASS
// ok      benchmark       0.002s

// Benchmark tests typically go in _test.go files and are named
// beginning with Benchmark. The testing runner executes each
// benchmark function several times, increasing b.N on each run
// until it collects a precise measurement.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

// $ go test -bench=.
// goos: linux
// goarch: amd64
// pkg: benchmark
// cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
// BenchmarkIntMin-4       1000000000               0.3306 ns/op
// PASS
// ok      benchmark       0.372s
