package authorization

import (
	"errors"
	"time"

	"github.com/S-Kiev/API-Go-Probando/modelo"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data *modelo.Login) (string, error) {
	claim := modelo.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Ezequiel",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(t string) (modelo.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &modelo.Claim{}, verifyFunction)
	if err != nil {
		return modelo.Claim{}, err
	}
	if !token.Valid {
		return modelo.Claim{}, errors.New("token no válido")
	}

	claim, ok := token.Claims.(*modelo.Claim)
	if !ok {
		return modelo.Claim{}, errors.New("no se pudo obtener los claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
