package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Encrypter struct {
	Cost int `json:"cost"`
}

func New(cost int) *Encrypter {
	return &Encrypter{
		Cost: cost,
	}
}

func (e *Encrypter) GenerateHash(src string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(src), e.Cost)
	return string(hashed), err
}

func (e *Encrypter) CompareHashAndPassword(raw, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw)) == nil
}
