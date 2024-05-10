package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpFile2.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", fromFile)
}
