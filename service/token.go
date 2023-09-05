package service

import (
	"time"
    // "fmt"
	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v4"
    "errors"
)

var tokenKey = []byte("aaa")

func GenerateToken(username string)(string, error){
    expireDuration, _ := time.ParseDuration("23h59m59s")
    expireTime := time.Now().Add(expireDuration)
    
    claims := jwt.StandardClaims{
            Audience: username,
            ExpiresAt: expireTime.Unix(),
            IssuedAt: time.Now().Unix(),
            Issuer: "demo", 
            NotBefore: time.Now().Unix(),
        }
    
    token , err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(tokenKey)
    // fmt.Println("token = " , token)
    
    return token, err
}

func ParseToken(token string)(string, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error) {
        return tokenKey, nil
    })

    if err != nil {
        return "", err
    }

    if tokenClaims == nil || !tokenClaims.Valid {
        return "", errors.New("invalid token")
    }

    claims := tokenClaims.Claims.(*jwt.StandardClaims)
    audience := claims.Audience
    return audience, nil
}