package encrypt

type Encryption interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	EncryptAndBase64([]byte) (string, error)
	DecryptFromBase64(string) ([]byte, error)
}
