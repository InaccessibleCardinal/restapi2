package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var myKey = []byte("Some Secret")

//var myKey = os.Getenv("TEST_JWT")

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, t)
}

//GenerateJWT generates a json web token
func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "ken"
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":3334", nil))
}

func main() {
	handleRequests()
}
