package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func getUsers(url string) ([]User, error) {
	var users []User
	res, err := http.Get(url)
	if err != nil {
		return users, err
	}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&users); err != nil {
		return users, err
	}
	return users, nil
}
