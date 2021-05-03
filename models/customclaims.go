package models

import "github.com/form3tech-oss/jwt-go"

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}
