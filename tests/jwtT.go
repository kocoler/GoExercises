package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

var key string = "mini"

func main() {
	/*token := jwt.New(jwt.SigningMethodES256)
	claims := make(jwt.MapClaims)
	claims[""]*/
	claims := &jwt.StandardClaims{
		Audience:  "2019",
		//IssuedAt: time.Now().Unix(),
		//ExpiresAt:time.Now().Add(time.Second* time.Duration(1000)).Unix(),
		Id:        "214",
		//NotBefore: int64(time.Now().Unix()+1000),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS256,claims)
	//key := "mini"
	//signKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte("key"))
	signedToken, err := token.SignedString([]byte("signKey"))
	if err != nil {
		fmt.Print("qwwq")
		log.Println(err)
	}
	fmt.Println(signedToken)
}
