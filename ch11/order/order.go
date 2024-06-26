package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/*
{
	"id": "12345",
	"date_ordered": "2020-05-01T13:01:02Z",
	"customer_id": "3",
	"items": [{"id": "xyz123", "name": "Thing 1"},{"id": "abc789","name": "Thing 2"}]
}
*/

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	str := `{
		"id": "12345",
		"date_ordered": "2020-05-01T13:01:02Z",
		"customer_id": "3",
		"items": [{"id": "xyz123", "name": "Thing 1"},{"id": "abc789","name": "Thing 2"}]
	}`
	data := []byte(str)

	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(o)
}
