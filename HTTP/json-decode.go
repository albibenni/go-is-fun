package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
