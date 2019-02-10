package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
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
	//fmt.Fprintf(w, t)
	data, err := ioutil.ReadFile("client/index.html")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404: " + http.StatusText(404)))
	}
	exp := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    "Token",
		Value:   t,
		Expires: exp,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
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
