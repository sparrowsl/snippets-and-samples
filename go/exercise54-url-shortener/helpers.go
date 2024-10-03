package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func generateShortURL(length int) string {
	var randomRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")

	if (length) <= 0 {
		length = 8
	}

	str := make([]rune, length)

	for i := range str {
		str[i] = randomRunes[rand.Intn(len(randomRunes))]
	}

	return string(str)
}

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
