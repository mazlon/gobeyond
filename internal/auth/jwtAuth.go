package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mazlon/gobeyond/config/config"
)

func main() {
	

}

func generateJWT() (string, error) {
	jwtSecret := config.GetTheEnv("JWTSECRET")
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true

}
