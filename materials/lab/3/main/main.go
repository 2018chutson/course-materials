// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>
// SHODAN_API_KEY=HBgJ7ImdwkOh15UvrUelRTthio7Q4aFN ./main webcam

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.AccountProfile()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Member: %t\nCredits:  %d\nDisplay Name:  %s\nCreated:  %s\n\n",
		info.Member,
		info.Credits,
		info.DisplayName,
		info.Created)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ", host.IPString, "====")
		h, _ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ", host.IPString, "====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}

	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}

}
