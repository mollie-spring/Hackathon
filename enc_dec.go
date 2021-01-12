import (
	"os"
	"/Users/molliezechlin/go/src/github.com/fentec-project/gofe/abe"
	"pubKey"
)

file, err := os.Open("pubkey")
n = file.Read(pubkey)
defer file.Close()

func encrypt(msg, msp)  {
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	return cipher
}

func decrypt(cipher, keys)  {
	dec, _ := a.Decrypt(cipher, keys, pubKey) // Decrypt the message
	return dec
}


