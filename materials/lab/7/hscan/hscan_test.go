// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "wordlist.txt")
	/* want := "foo"
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} */

}

// help with benchmark from here: https://flaviocopes.com/golang-measure-time/
func BenchmarkGenHashMaps(b *testing.B) {
	f := "tiny.txt"
	for i := 0; i < b.N; i++ {
		GenHashMaps(f)
	}
}

// RESULTS OF BENCHMARK: (both using the "tiny.txt" password file, which contains 303,872 passwords)
//
//		Without concurrency:
// $ go test -bench=.
// [+] Password found (MD5): Nickelback4life
// goos: windows
// goarch: amd64
// pkg: hscan/hscan
// cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
// BenchmarkGenHashMaps-8                 2         644625100 ns/op
// PASS
//
//		time per password: 644,625,100 / 303,872 = 2,121.37 ns  (not sure if that is the correct way to find that)
//
//		With Concurrency:
// $ go test -bench=.
// [+] Password found (MD5): Nickelback4life
// goos: windows
// goarch: amd64
// pkg: hscan/hscan
// cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
// BenchmarkGenHashMaps-8                 3         444834000 ns/op
// PASS
// ok      hscan/hscan     3.476s
//
//		time per password: 444,834,000 / 303,872 = 1,463.89 ns
