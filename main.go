package main

import (
	"context"
	"fmt"
	"net"
	"sync"
)

const (
	concurrency    = 10
	alphabet       = "abcdefghijklmnopqrstuvwxyz"
	consonants     = "bcdfghjklmnpqrstvwxz"
	niceConsonants = "bcdfghjklmnpqrstvxz"
	vowels         = "aeiouy"
	niceVowels     = "aeiou"
	// numbers     = "0123456789"
	numbers = ""
)

func main() {
	data := twowords("moto", "eu")

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
				d := string(i) + string(j) + "." + tld
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
	}()

	return dc
}

func short(word string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, tld := range tlds {
			dc <- word + "." + tld
		}
		close(dc)
	}()

	return dc
}

func short2(word string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range alphabet + numbers {
			for _, j := range alphabet + numbers {
				for _, tld := range tlds {
					dc <- word + string(i) + string(j) + "." + tld
					// dc <- string(i) + string(j) + word + "." + tld
					// dc <- string(i) + word + string(j) + "." + tld
				}
			}
		}
		close(dc)
	}()

	return dc
}

func short3(word string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, tld := range tlds {
			for _, i := range niceVowels {
				for _, j := range niceConsonants {
					for _, k := range niceVowels {
						dc <- fmt.Sprintf("%s%s%s%s.%s", word, string(i), string(j), string(k), tld)
					}
				}
			}
			for _, i := range niceConsonants {
				for _, j := range niceVowels {
					for _, k := range niceConsonants {
						dc <- fmt.Sprintf("%s%s%s%s.%s", word, string(i), string(j), string(k), tld)
					}
				}
			}
		}
		close(dc)
	}()

	return dc
}

func short4(suffix string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, tld := range tlds {
			for _, i := range niceVowels {
				for _, j := range niceConsonants {
					for _, k := range niceConsonants {
						for _, l := range niceVowels {
							dc <- fmt.Sprintf("%s%s%s%s%s.%s", string(i), string(j), string(k), string(l), suffix, tld)
						}
					}
				}
			}
		}
		close(dc)
	}()

	return dc
}

func twowords(word string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, i := range letters3() {
			for _, tld := range tlds {
				dc <- word + string(i) + "." + tld
				dc <- string(i) + word + "." + tld
			}
		}
		close(dc)
	}()

	return dc
}

func oneword(tlds ...string) <-chan string {
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

func prefixSuffix(word string, tlds ...string) <-chan string {
	dc := make(chan string)

	go func() {
		for _, tld := range tlds {
			for _, i := range prefix() {
				dc <- string(i) + word + "." + tld
			}
			for _, i := range suffix() {

				dc <- word + string(i) + "." + tld
			}
		}
		close(dc)
	}()

	return dc
}
