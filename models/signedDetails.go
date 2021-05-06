package models

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email     string
	Full_name string
	School    string
	Standard  int
	Uid       string
	jwt.StandardClaims
}
