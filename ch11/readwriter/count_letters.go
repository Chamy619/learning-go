package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := make(map[string]int)

	for {
		_, err := r.Read(buf)
		for _, b := range buf {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err != io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	file, err := os.Open("text")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	cl, err := countLetters(file)
	if err != nil {
		log.Fatal(err)
	}

	for l, c := range cl {
		fmt.Printf("%s: %d\n", l, c)
	}
}
