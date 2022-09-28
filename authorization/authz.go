package authorization

import (
	"backChannel/config"
	"github.com/golang-jwt/jwt"
	"log"
)

type Claim struct {
	ResourceAccess map[string]ClaimClient `json:"resource_access"`
}

type ClaimClient struct {
	Roles []string `json:"roles"`
}

func AuthorizeUser(tokenString string) bool {
	log.Println(tokenString)
	claims := jwt.MapClaims{}
	//claims := Claim{}
	_, err := jwt.ParseWithClaims(tokenString, claims, nil)
	if err != nil {
		log.Println("error decoding JWT", err)
	}
	mp := claims["resource_access"]

	for key, _ := range mp.(map[string]interface{}) {
		if key == config.ClientId {
			return true
		}
	}

	return false
}
