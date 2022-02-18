// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
//
// Carly Hutson
// Lab 2
// 2/11/22

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int // notice the capitalization here. access limited!
var closedports []int

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(ps []int) int {

	ports := make(chan int, 3000) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := range ps {
			ports <- ps[i]
		}
	}()

	for i := 0; i < len(ps); i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closedports = append(closedports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	// I put in a check if the user is playing World of Warcraft or not, since it's my boyfriend's favorite game. :)
	playingWow := false
	fmt.Print("---------OPEN PORTS:------------\n")
	for _, port := range openports {
		fmt.Printf("%d\n", port)
		if port == 3724 {
			playingWow = true
		}
	}
	fmt.Printf("-------The other %d ports are closed.---------\n", len(closedports))
	if playingWow {
		fmt.Print("You are currently playing WoW!\n")
	} else {
		fmt.Print("You are not currently playing WoW! :(\n")
	}
	return len(openports) + len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed);
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
