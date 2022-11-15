package main

import (
	"fmt"
	"time"

	//"github.com/matryer/moq/generate"
	"github.com/dgrijalva/jwt-go"
)

var mysecretphrase = []byte("my-super-secret-Token")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = "93670"
	claims["Name"] = "Roshan kumar Jalan"
	claims["PhNo"] = "7667278529"
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(25 * time.Second).Unix()

	tokenString, err := token.SignedString(mysecretphrase)
	return tokenString, err
}

// func verifyJWT(tokenString string) bool {
// 	theToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("Invalid")
// 		}
// 		return mysecretphrase, nil
// 	})
// 	if err != nil {
// 		return false
// 	}
// 	return theToken.Valid
// }

func main() {
	fmt.Println("Jwt Token created successfully")
	tokenString, _ := generateJWT()
	fmt.Println(tokenString)
	//ok := verifyJWT(tokenString)
	//fmt.Println(ok)
}
