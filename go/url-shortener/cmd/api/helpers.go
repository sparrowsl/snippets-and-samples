package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// Generates a short random URL with a default length of 8 characters.
func generateShortURL() string {
	var randomRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")

	str := make([]rune, 8)

	for i := range str {
		str[i] = randomRunes[rand.Intn(len(randomRunes))]
	}

	return string(str)
}

// Sends a JSON response back to the client.
func toJSON(writer http.ResponseWriter, status int, data map[string]any) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(result)

	return nil
}
