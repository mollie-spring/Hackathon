package main

import (
	"os"
	"github.com/fentec-project/gofe/abe"
)



func encrypt(msg string, msp *MSP)  {
	file, err := os.Open("keys/pubkey")
	n = file.Read(pubKey *FAMEPubKey)
	defer file.Close()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	return cipher
}

func decrypt(cipher *FAMECipher, keys *FAMESecKey)  {
	file, err := os.Open("keys/pubkey")
	n = file.Read(pubkey)
	defer file.Close()
	dec, _ := a.Decrypt(cipher, keys, pubKey) // Decrypt the message
	return dec
}


