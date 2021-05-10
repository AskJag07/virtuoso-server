package models

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Full_name string
	Session   int
	Uid       string
	Admin     bool
	jwt.StandardClaims
}
