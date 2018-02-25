package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WhosInSpace : comment
type WhosInSpace struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	People  []*Person
}

// Person : comment
type Person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

func main() {
	url := "http://api.open-notify.org/astros.json"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Println(err)
		os.Exit(1)
	}

	var wis WhosInSpace

	if err := json.NewDecoder(resp.Body).Decode(&wis); err != nil {
		resp.Body.Close()
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("number of people in space %d\n", wis.Number)
	for _, p := range wis.People {
		fmt.Println(p.Craft, p.Name)
	}
}
