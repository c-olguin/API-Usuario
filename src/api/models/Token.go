package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)


//Estructura del token
type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}