package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SpacePeople struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

type SpaceData struct {
	People  []SpacePeople `json:"people"`
	Number  int           `json:"number"`
	Message string        `json:"message"`
}

func main() {
	response, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	spaceData := SpaceData{}
	if err := json.Unmarshal(data, &spaceData); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("There are %d people in space right now:\n\n", spaceData.Number)

	fmt.Printf("%-23s Craft\n", "Name")
	fmt.Println(strings.Repeat("-", 30))
	for _, person := range spaceData.People {
		fmt.Printf("%-23s %s\n", person.Name, person.Craft)
	}

}
