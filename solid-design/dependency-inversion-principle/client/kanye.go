package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type kanye struct{}

type quotes struct {
	Quote string
}

func NewKanye() *kanye {
	return &kanye{}
}

func (k *kanye) GetQuote() (string, error) {
	resp, err := http.Get("https://api.kanye.rest")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var quotes quotes
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return "", err
	}

	return quotes.Quote, nil
}
