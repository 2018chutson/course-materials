package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {
	s := []int{22, 80, 135, 1725, 3724, 50520}
	got := PortScanner(s) // Currently function returns only number of open ports
	want := len(s)        // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	//consider what would happen if you parameterize the portscanner address and ports to scan

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTotalPortsScanned(t *testing.T) {
	s := []int{22, 80, 135, 1725, 3724, 50520}
	got := PortScanner(s)
	want := len(s) * 2 // default value; consider what would happen if you parameterize the portscanner ports to scan

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
