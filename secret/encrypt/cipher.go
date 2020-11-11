package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"io"
)

//ReturnEncryptStream function accepts a key string and input value bytes and returns a cipher stream
func ReturnEncryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	return cipher.NewCFBEncrypter(block, iv), err
}

//ReturnDecryptStream function accepts a key string and input value bytes and returns a cipher stream
func ReturnDecryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	return cipher.NewCFBDecrypter(block, iv), err
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}

//EncryptorWriter returns a streamwriter which encrypts the incoming content
func EncryptorWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, iv)
	stream, err := ReturnEncryptStream(key, iv)
	_, err = w.Write(iv)
	return &cipher.StreamWriter{S: stream, W: w}, err
}

//DecryptorReader returns a stream reader which decrypts the incoming content
func DecryptorReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if err != nil || n != len(iv) {
		return nil, errors.New("Encrypt: Unable to read IV")
	}
	stream, err := ReturnDecryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, err
}
