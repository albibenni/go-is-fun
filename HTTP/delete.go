package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest(http.MethodDelete, fullURL, nil)
	if err != nil {
		return err
	}

	// set request headers
	req.Header.Set("X-API-Key", apiKey)

	// create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("request to delete user unsuccessful")
	}

	return nil
}
