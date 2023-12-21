package jwt

import (
	"crypto/rsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Spargwy/just-to-do-it/pkg/auth/model"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

const (
	privateKeyPath = "tasker.rsa"
	publicKeyPath  = "tasker.rsa.pub"
)

type Authenticator struct {
	signBytes []byte
	signKey   *rsa.PrivateKey

	verifyBytes []byte
	verifyKey   *rsa.PublicKey
}

func New(path string) *Authenticator {
	if len(path) < 1 {
		logrus.Fatal("JWT key path is empty")
	}

	signBytes, err := os.ReadFile(filepath.Join(path, privateKeyPath))
	if err != nil {
		logrus.Panicf("%+v", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logrus.Panicf("%+v", err)
	}
	verifyBytes, err := os.ReadFile(filepath.Join(path, publicKeyPath))
	if err != nil {
		logrus.Panicf("%+v", err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logrus.Panicf("%+v", err)
	}

	return &Authenticator{
		signBytes:   signBytes,
		signKey:     signKey,
		verifyBytes: verifyBytes,
		verifyKey:   verifyKey,
	}
}

func (j *Authenticator) Generate(claims *model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(j.signKey)
	return tokenString, err
}

func (j *Authenticator) Parse(t string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(t, &model.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("failed to parse token")
	}

	return claims, nil
}
