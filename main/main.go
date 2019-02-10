package main

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

var myKey = []byte("Some Secret")

func main() {
	handleRequests()
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "homepage hit")

}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if len(id) > 0 {

		handleUser(w, r, id)

	} else {

		usersBytes := getUsersFromProxy()
		u, err := json.Marshal(usersBytes)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Server error"))
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(u)
	}

}

func handleUser(w http.ResponseWriter, r *http.Request, id string) {

	userBytes := getUserFromProxy(id)
	u, err := json.Marshal(userBytes)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Server error"))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(u)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// log.Println(r.Header["Token"])
			//w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		if r.Header["Token"] != nil {
			log.Println("Header: ", r.Header["Token"])
			token, err := jwt.Parse(
				r.Header["Token"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("error")
					}
					return myKey, nil
				})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}

		} else {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
		}

	})
}

func handleRequests() {
	http.HandleFunc("/", handleRoutes)
	http.Handle("/users/", isAuthorized(handleUsers))
	log.Fatal(http.ListenAndServe(":3333", nil))
}
