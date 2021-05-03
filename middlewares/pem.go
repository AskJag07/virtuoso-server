package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/form3tech-oss/jwt-go"

	"github.com/AskJag07/virtuoso-server/models"
)

func GetPemCert(token *jwt.Token) (string, error) {

	cert := ""
	resp, err := http.Get("https://virtuoso.eu.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = models.Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil

}
