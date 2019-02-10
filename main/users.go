package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//User base user object
type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Website string  `json:"website"`
	Address Address `json:"address"`
	Company Company `json:"company"`
}

//Address user address object
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

//Company user company object
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

const url = "https://jsonplaceholder.typicode.com/users/"

func getUsersFromProxy() []User {
	response, getErr := http.Get(url)
	if getErr != nil {
		log.Fatalln(getErr)
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalln(getErr)
	}
	return makeUsers(body)
}

func makeUsers(body []byte) []User {
	u := []User{}
	err := json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	return u
}

func getUserFromProxy(id string) User {
	userURL := url + string(id)
	r, getErr := http.Get(userURL)
	if getErr != nil {
		log.Fatalln(getErr)
	}
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		panic(readErr)
	}
	return makeUser(body)
}

func makeUser(body []byte) User {
	u := User{}
	err := json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	return u
}
