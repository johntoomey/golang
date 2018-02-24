package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	resp, _ := client.Get(url)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}

	var wis WhosInSpace

	if err := json.NewDecoder(resp.Body).Decode(&wis); err != nil {
		resp.Body.Close()
	}

	fmt.Printf("people in space %d\n", wis.Number)
	for _, p := range wis.People {
		fmt.Println(p.Name)
	}

}
