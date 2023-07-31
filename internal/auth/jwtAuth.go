package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id string) (string, error) {
	token_lifespan := 24
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWTSECRET")))
}

func TokenVerifier(tokenString string) (string, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTSECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {

		return "", errors.New("couldn't parse the token")
	}
	userID, ok := claims[`user_id`].(string)
	if !ok {
		err = errors.New("couldn't extract user_id from token")
		return "", err
	}
	return userID, nil
}

//TODO unit test that generates token and feeds it to tokenverifier and checks if it works

// func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {

// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		if request.Header["Token"] != nil {

// 		}
// 		token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
//             _, ok := token.Method.(*jwt.SigningMethodECDSA)
//             if !ok {
//                writer.WriteHeader(http.StatusUnauthorized)
//                _, err := writer.Write([]byte("You're Unauthorized!"))
//                if err != nil {
//                   return nil, err

//                }
//             }
//             return "", nil

//          })
// 	})
// }
