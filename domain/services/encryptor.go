package services

type Encryptor interface {
	Encrypt(text string) (string, error)
	Check(hashedText string, text string) error
}
