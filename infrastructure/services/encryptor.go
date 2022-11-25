package services

import "golang.org/x/crypto/bcrypt"

type Encryptor struct {
}

func NewEncryptor() Encryptor {
	return Encryptor{}
}

func (e Encryptor) Encrypt(text string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	if err != nil {
		return "", nil
	}
	return string(encrypted), nil
}

func (e Encryptor) Check(hashedTest string, text string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedTest), []byte(text))
}
