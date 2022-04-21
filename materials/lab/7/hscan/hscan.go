package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if len(hash) != len(sourceHash) {
			hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		} else {
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GenHashMaps(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	md5lookup := make(map[string]string)
	shalookup := make(map[string]string)
	// needed this RWMutex to lock/unlock maps to solve concurrency map write errors
	// found here: https://stackoverflow.com/questions/45585589/golang-fatal-error-concurrent-map-read-and-map-write
	md5lookupMutex := sync.RWMutex{}
	shalookupMutex := sync.RWMutex{}

	for scanner.Scan() {
		password := scanner.Text()

		go func() {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			md5lookupMutex.Lock()
			md5lookup[hash] = password
			md5lookupMutex.Unlock()
		}()

		go func() {
			shahash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			shalookupMutex.Lock()
			shalookup[shahash] = password
			shalookupMutex.Unlock()
		}()

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}
