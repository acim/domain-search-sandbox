package main

import (
	"context"
	"fmt"
	"net"
	"sync"
)

const (
	concurrency = 10
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
	// numbers     = "0123456789"
	numbers = ""
)

func main() {
	data := oneword([]string{"sh"})

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go lookup(i, &wg, data)
	}

	wg.Wait()
}

func lookup(i int, wg *sync.WaitGroup, data <-chan string) {
	defer wg.Done()

	for {
		d, ok := <-data
		if !ok {
			return
		}

		_, err := net.DefaultResolver.LookupIPAddr(context.Background(), "www."+d)
		if err != nil {
			_, err = net.DefaultResolver.LookupIPAddr(context.Background(), d)
			if err != nil {
				fmt.Println(d)
			}
		}
	}
}

func generator4(tld string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet {
			for _, j := range alphabet {
				for _, k := range alphabet {
					for _, l := range alphabet {
						d := string(i) + string(j) + string(k) + string(l) + tld
						dc <- d
					}
				}
			}
		}
		close(dc)
	}()

	return dc
}

func generator2(tld string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet + numbers {
			for _, j := range alphabet + numbers {
				d := string(i) + string(j) + tld
				dc <- d
			}
		}
		close(dc)
	}()

	return dc
}

func generator3(tlds []string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet + numbers {
			for _, j := range alphabet + numbers {
				for _, k := range alphabet + numbers {
					for _, tld := range tlds {
						d := string(i) + string(j) + string(k) + "." + tld
						dc <- d
					}
				}
			}
		}
		close(dc)
	}()

	return dc
}

func short(word string, tlds []string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet + numbers {
			for _, tld := range tlds {
				dc <- word + string(i) + "." + tld
				dc <- string(i) + word + "." + tld
			}
		}
		close(dc)
	}()

	return dc
}

func short2(word string, tld string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet + numbers {
			for _, j := range alphabet + numbers {
				dc <- word + string(i) + string(j) + tld
				dc <- string(i) + string(j) + word + tld
				dc <- string(i) + word + string(j) + tld
			}
		}
		close(dc)
	}()

	return dc
}

func twowords(word string, tlds []string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range letters5() {
			for _, tld := range tlds {
				dc <- word + string(i) + "." + tld
				dc <- string(i) + word + "." + tld
			}
		}
		close(dc)
	}()

	return dc
}

func oneword(tlds []string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range letters3() {
			for _, tld := range tlds {
				dc <- string(i) + "." + tld
			}
		}
		close(dc)
	}()

	return dc
}
