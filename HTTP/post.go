package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// ?    // encode our User as json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// decode the json data from the response
	// into a new User struct
	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
