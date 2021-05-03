package models

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}
